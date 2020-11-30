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
		return ddderr.NewInvalidFormat("command", cmd.Id())
	}

	preCrawlUrl, err := movie.NewCrawlUrl(command.CrawlUrl)
	if err != nil {
		return err
	}

	crawled, err := h.crawler.Crawl(ctx, *preCrawlUrl)
	if err != nil {
		return err
	}

	return h.createMovie(ctx, command, crawled)
}

func (h CreateByCrawlCommandHandler) createMovie(ctx context.Context, command CreateByCrawlCommand,
	crawled *movie.CrawledMovie) error {
	var resultErr *multierror.Error
	movieId, err := movie.NewMovieId(command.MovieId)
	if err != nil {
		return err
	}
	userId, err := movie.NewUserId(command.UserId)
	if err != nil {
		return err
	}
	displayName, err := movie.NewDisplayName(crawled.Title)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	description, err := movie.NewDescription(crawled.Plot)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	year, err := movie.NewYear(uint32(crawled.Year))
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	picture, err := movie.NewPicture(crawled.Poster)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	crawlUrl, err := movie.NewCrawlUrl(crawled.CrawlUrl)
	if err != nil {
		return err
	}

	if resultErr != nil && resultErr.ErrorOrNil() != nil {
		return resultErr
	}
	return h.creator.Create(ctx, CreatorArgs{
		Id:          *movieId,
		UserId:      *userId,
		DisplayName: *displayName,
		Description: *description,
		Year:        *year,
		Picture:     *picture,
		WatchUrl:    movie.WatchUrl{},
		CrawlUrl:    *crawlUrl,
	})
}
