package movie

import (
	"context"
)

// CrawledMovie film crawled from an external system, contains primitives
//	@DTO
type CrawledMovie struct {
	Title  string
	Plot   string
	Year   int
	Poster string
}

// MovieCrawler fetch movie data from an external system by URL
//	@InfrastructureService
//	@Port
type MovieCrawler interface {
	Crawl(ctx context.Context, url CrawlUrl) (*CrawledMovie, error)
}
