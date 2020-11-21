package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/neutrinocorp/ddderr"

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

func (h FindQueryHandler) Invoke(ctx context.Context, q domain.Query) (interface{}, error) {
	query, ok := q.(FindQuery)
	if !ok {
		return nil, ddderr.NewInvalidFormat("query", "find query")
	}

	movieId, err := movie.NewMovieId(query.MovieId)
	if err != nil {
		return nil, err
	}
	return h.finder.Find(ctx, *movieId)
}
