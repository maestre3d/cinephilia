package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
	"github.com/neutrinocorp/ddderr"
)

// FindQueryHandler request a searching of an actor by unique identifier
//	@QueryHandler
type FindQueryHandler struct {
	finder *Finder
}

func NewFindQueryHandler(finder *Finder) *FindQueryHandler {
	return &FindQueryHandler{finder: finder}
}

func (h FindQueryHandler) Invoke(ctx context.Context, query domain.Query) (interface{}, error) {
	q, ok := query.(FindQuery)
	if !ok {
		return nil, ddderr.NewInvalidFormat("query", query.Id())
	}

	id, err := actor.NewId(q.ActorId)
	if err != nil {
		return nil, err
	}

	act, err := h.finder.Find(ctx, *id)
	if err != nil {
		return nil, err
	}
	return newResponse(*act), nil
}
