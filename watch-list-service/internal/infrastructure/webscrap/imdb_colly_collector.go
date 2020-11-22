package webscrap

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func NewCollyImdbCollector() (*colly.Collector, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
		// colly.Async(true),
	)
	err := c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "www.imdb.com/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
	})

	return c, err
}
