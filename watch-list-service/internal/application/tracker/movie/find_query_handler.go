package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/neutrinocorp/ddderr"
)

// FindQueryHandler receives a FindQuery and search for a movie by Id
//	@QueryHandler
type FindQueryHandler struct {
	finder *Finder
}

func NewFindQueryHandler(bus domain.QueryBus, finder *Finder) (*FindQueryHandler, error) {
	h := &FindQueryHandler{finder: finder}
	if err := bus.RegisterHandler(FindQuery{}, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (h FindQueryHandler) Invoke(ctx context.Context, q domain.Query) (interface{}, error) {
	query, ok := q.(FindQuery)
	if !ok {
		return nil, ddderr.NewInvalidFormat("query", q.Id())
	}

	movieId, err := movie.NewMovieId(query.MovieId)
	if err != nil {
		return nil, err
	}
	return h.finder.Find(ctx, *movieId)
}
