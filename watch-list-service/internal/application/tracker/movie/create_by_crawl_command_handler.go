package movie

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/neutrinocorp/ddderr"
)

// CreateByCrawlCommandHandler receives CreateByCrawlCommand(s), creates a Movie by fetching an external url
//	@CommandHandler
//	@Async
type CreateByCrawlCommandHandler struct {
	creator *Creator
	crawler movie.MovieCrawler
}

func NewCreateByCrawlCommandHandler(creator *Creator, crawler movie.MovieCrawler) *CreateByCrawlCommandHandler {
	return &CreateByCrawlCommandHandler{creator: creator, crawler: crawler}
}

func (h CreateByCrawlCommandHandler) Invoke(ctx context.Context, cmd domain.Command) error {
	command, ok := cmd.(CreateByCrawlCommand)
	if !ok {
		return ddderr.NewInvalidFormat("command", "create_by_crawl command")
	}
	movieId, err := movie.NewMovieId(command.MovieId)
	if err != nil {
		return err
	}
	userId, err := movie.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	crawlUrl, err := movie.NewCrawlUrl(command.CrawlUrl)
	if err != nil {
		return err
	}
	movCrawled, err := h.crawler.Crawl(ctx, *crawlUrl)
	if err != nil {
		return err
	}

	var resultErr *multierror.Error
	displayName, err := movie.NewDisplayName(movCrawled.Title)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	description, err := movie.NewDescription(movCrawled.Plot)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	year, err := movie.NewYear(uint32(movCrawled.Year))
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	picture, err := movie.NewPicture(movCrawled.Poster)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}

	if resultErr != nil && resultErr.ErrorOrNil() != nil {
		return resultErr
	}
	return h.creator.Create(ctx, CreatorArgs{
		Id:          *movieId,
		UserId:      *userId,
		CategoryId:  movie.CategoryId{},
		DirectorId:  movie.DirectorId{},
		DisplayName: *displayName,
		Description: *description,
		Year:        *year,
		Picture:     *picture,
		WatchUrl:    movie.WatchUrl{},
		CrawlUrl:    *crawlUrl,
	})
}
