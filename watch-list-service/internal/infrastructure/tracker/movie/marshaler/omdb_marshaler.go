package marshaler

import (
	"strconv"
	"time"

	"github.com/eefret/gomdb"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

func OmdbToAggregate(movieOmdb *gomdb.MovieResult) *movie.CrawledMovie {
	year, err := strconv.ParseInt(movieOmdb.Year, 10, 64)
	if err != nil {
		year = int64(time.Now().UTC().Year())
	}
	return &movie.CrawledMovie{
		Title:  movieOmdb.Title,
		Plot:   movieOmdb.Plot,
		Year:   int(year),
		Poster: movieOmdb.Poster,
	}
}
