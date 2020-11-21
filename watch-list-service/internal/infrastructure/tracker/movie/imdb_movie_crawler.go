package movie

import (
	"context"
	"strings"

	"github.com/neutrinocorp/ddderr"

	"github.com/eefret/gomdb"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/marshaler"
)

//	@InfrastructureService
//	@Adapter
type ImdbMovieCrawler struct {
	api *gomdb.OmdbApi
}

func NewImdbMovieCrawler(api *gomdb.OmdbApi) *ImdbMovieCrawler {
	return &ImdbMovieCrawler{api: api}
}

func (i ImdbMovieCrawler) Crawl(_ context.Context, url movie.CrawlUrl) (*movie.CrawledMovie, error) {
	movRes, err := i.api.MovieByImdbID(i.getIdFromUrl(url))
	if err != nil {
		return nil, ddderr.NewInfrastructure(err, "api error")
	}

	return marshaler.ImdbToAggregate(movRes), nil
}

func (i ImdbMovieCrawler) getIdFromUrl(url movie.CrawlUrl) string {
	titleSlice := strings.Split(url.Value(), "/title/")
	if len(titleSlice) < 1 {
		return ""
	}

	imdbIdSlice := strings.Split(titleSlice[1], "/")
	if len(imdbIdSlice) == 0 {
		return ""
	}

	return imdbIdSlice[0]
}
