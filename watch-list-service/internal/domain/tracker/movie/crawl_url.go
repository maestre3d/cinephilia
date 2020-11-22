package movie

import (
	"strings"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// CrawlUrl film external data URL, Cinephilia will crawl asynchronously when requested
//
//	Only IMDb sites are allowed for fetching/crawling by the moment
//	@ValueObject
type CrawlUrl struct {
	url domain.Url
}

func NewCrawlUrl(value string) (*CrawlUrl, error) {
	crawlUrl := new(CrawlUrl)
	if err := crawlUrl.ensureIsImdb(value); value != "" && err != nil {
		return nil, err
	}

	url, err := domain.NewUrl("crawl_url", value)
	if err != nil {
		return nil, err
	}

	crawlUrl.url = *url
	return crawlUrl, nil
}

func (u CrawlUrl) IsEmpty() bool {
	return u.Value() == ""
}

func (u CrawlUrl) ensureIsImdb(value string) error {
	//	rules
	//	a.	value == imdb valid domain
	isImdb := strings.HasPrefix(value, "https://www.imdb.com") || strings.HasPrefix(value,
		"http://www.imdb.com") || strings.HasPrefix(value, "https://imdb.com") ||
		strings.HasPrefix(value, "http://imdb.com")
	if !isImdb {
		return CrawlUrlIsNotAvailable
	}

	return nil
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
