package movie

import (
	"context"
	"strings"

	"github.com/eefret/gomdb"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/marshaler"
	"github.com/neutrinocorp/ddderr"
)

//	@InfrastructureService
//	@Adapter
type OmdbMovieCrawler struct {
	api *gomdb.OmdbApi
}

func NewOmdbMovieCrawler(api *gomdb.OmdbApi) *OmdbMovieCrawler {
	return &OmdbMovieCrawler{api: api}
}

func (i OmdbMovieCrawler) Crawl(_ context.Context, url movie.CrawlUrl) (*movie.CrawledMovie, error) {
	movRes, err := i.api.MovieByImdbID(i.getIdFromUrl(url))
	if err != nil {
		return nil, ddderr.NewInfrastructure(err, "api error")
	}

	return marshaler.OmdbToAggregate(movRes), nil
}

func (i OmdbMovieCrawler) getIdFromUrl(url movie.CrawlUrl) string {
	titleSlice := strings.Split(url.Value(), "/title/")
	if len(titleSlice) < 1 {
		return ""
	}

	OmdbIdSlice := strings.Split(titleSlice[1], "/")
	if len(OmdbIdSlice) == 0 {
		return ""
	}

	return OmdbIdSlice[0]
}
