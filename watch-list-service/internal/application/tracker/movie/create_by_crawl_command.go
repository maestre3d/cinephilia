package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// CreateByCrawlCommand requests a movie creation by crawling technique
//	@Command
//	@Async
//	@DTO
type CreateByCrawlCommand struct {
	MovieId  string
	UserId   string
	CrawlUrl string
}

func (c CreateByCrawlCommand) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.movie.create_by_crawl"
}
