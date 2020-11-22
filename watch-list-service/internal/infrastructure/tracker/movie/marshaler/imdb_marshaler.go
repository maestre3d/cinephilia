package marshaler

import (
	"strconv"
	"strings"
	"time"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

func ImdbToAggregate(property, content string, movieRef *movie.CrawledMovie) {
	switch property {
	case "og:title":
		setTitleAndYear(content, movieRef)
	case "og:description":
		movieRef.Plot = content
	case "og:url":
		movieRef.CrawlUrl = content
	case "pageId":
		movieRef.ExternalId = content
	case "og:image":
		movieRef.Poster = content
	}
}

func setTitleAndYear(title string, movieRef *movie.CrawledMovie) {
	titleSlice := strings.Split(title, " (")
	if len(titleSlice) == 0 {
		return
	}
	movieRef.Title = titleSlice[0]

	yearSlice := strings.Split(titleSlice[1], ")")
	if len(yearSlice) == 0 {
		return
	}
	year, err := strconv.ParseInt(yearSlice[0], 10, 64)
	if err != nil {
		year = int64(time.Now().UTC().Year())
	}
	movieRef.Year = int(year)
}
