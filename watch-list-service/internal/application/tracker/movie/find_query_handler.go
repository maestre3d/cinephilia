package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

// FindQueryHandler receives a FindQuery and search for a movie by Id
//	@QueryHandler
type FindQueryHandler struct {
	finder *Finder
}

func NewFindQueryHandler(finder *Finder) *FindQueryHandler {
	return &FindQueryHandler{finder: finder}
}

func (h FindQueryHandler) Invoke(ctx context.Context, query FindQuery) (*MovieResponse, error) {
	movieId, err := movie.NewMovieId(query.Id)
	if err != nil {
		return nil, err
	}
	return h.finder.Find(ctx, *movieId)
}
