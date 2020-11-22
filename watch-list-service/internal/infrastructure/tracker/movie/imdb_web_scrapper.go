package movie

import (
	"github.com/gocolly/colly/v2"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/marshaler"
)

type ImdbWebScrapper struct {
	collector *colly.Collector
}

func NewImdbWebScrapper(collector *colly.Collector) *ImdbWebScrapper {
	return &ImdbWebScrapper{collector: collector}
}

func (s ImdbWebScrapper) Scrap(scrapUrl string) (crawledMovie *movie.CrawledMovie, err error) {
	crawledMovie = new(movie.CrawledMovie)
	s.collector.OnHTML("meta", func(e *colly.HTMLElement) {
		marshaler.ImdbToAggregate(e.Attr("property"), e.Attr("content"), crawledMovie)
	})
	s.collector.OnError(func(r *colly.Response, errC error) {
		err = errC
	})
	err = s.collector.Visit(scrapUrl)

	return crawledMovie, err
}
