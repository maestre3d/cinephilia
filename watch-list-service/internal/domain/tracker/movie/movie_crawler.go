package movie

import (
	"context"
)

// MovieCrawler fetch movie data from an external system by URL
//	@InfrastructureService
//	@Port
type MovieCrawler interface {
	Crawl(ctx context.Context, url CrawlUrl) (*Movie, error)
}
