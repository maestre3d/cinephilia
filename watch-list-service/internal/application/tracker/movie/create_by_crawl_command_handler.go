package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

// CreateByCrawlCommandHandler receives CreateByCrawlCommand(s), creates a Movie by fetching an external url
//	@CommandHandler
//	@Async
type CreateByCrawlCommandHandler struct {
	creator Creator
	crawler movie.MovieCrawler
}

func NewCreateByCrawlCommandHandler(creator Creator, crawler movie.MovieCrawler) *CreateByCrawlCommandHandler {
	return &CreateByCrawlCommandHandler{creator: creator, crawler: crawler}
}

func (h CreateByCrawlCommandHandler) Invoke(ctx context.Context, command CreateByCrawlCommand) error {
	crawlUrl, err := movie.NewCrawlUrl(command.CrawlUrl)
	if err != nil {
		return err
	}
	mov, err := h.crawler.Crawl(ctx, *crawlUrl)
	if err != nil {
		return err
	}

	return h.creator.Create(ctx, CreatorArgs{
		Id:          mov.Id(),
		UserId:      mov.User(),
		CategoryId:  mov.Category(),
		DirectorId:  mov.Director(),
		DisplayName: mov.DisplayName(),
		Description: mov.Description(),
		Year:        mov.Year(),
		Picture:     mov.Picture(),
		WatchUrl:    mov.WatchUrl(),
		CrawlUrl:    mov.CrawlUrl(),
	})
}
