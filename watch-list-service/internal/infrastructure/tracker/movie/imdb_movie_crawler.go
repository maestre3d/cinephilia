package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

//	@InfrastructureService
//	@Adapter
type ImdbMovieCrawler struct {
	scrapper *ImdbWebScrapper
}

func NewImdbMovieCrawler(scrapper *ImdbWebScrapper) *ImdbMovieCrawler {
	return &ImdbMovieCrawler{scrapper: scrapper}
}

func (i ImdbMovieCrawler) Crawl(_ context.Context, url movie.CrawlUrl) (*movie.CrawledMovie, error) {
	return i.scrapper.Scrap(url.Value())
}
