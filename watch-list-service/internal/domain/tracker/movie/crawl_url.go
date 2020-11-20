package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// CrawlUrl film external data URL, Cinephilia will crawl asynchronously when requested
//
//	Only IMDb sites are allowed for fetching/crawling by the moment
//	@ValueObject
type CrawlUrl struct {
	url domain.Url
}

func NewCrawlUrl(value string) (*CrawlUrl, error) {
	url, err := domain.NewUrl("crawl_url", value)
	if err != nil {
		return nil, err
	}

	return &CrawlUrl{url: *url}, nil
}

func (u CrawlUrl) IsEmpty() bool {
	return u.Value() == ""
}

func (u CrawlUrl) Value() string {
	return u.url.Value()
}

func (u CrawlUrl) NullableValue() *string {
	if u.Value() == "" {
		return nil
	}

	val := u.Value()
	return &val
}
