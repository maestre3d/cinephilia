package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
)

// Finder search for an actor by its unique identifier
//	@DomainService
type Finder struct {
	repository actor.Repository
}

func newFinder(repository actor.Repository) *Finder {
	return &Finder{repository: repository}
}

func (f Finder) Find(ctx context.Context, id actor.Id) (*actor.Actor, error) {
	act, err := f.repository.Search(ctx, id)
	if err != nil {
		return nil, err
	} else if act == nil {
		return nil, actor.NotFound
	}

	return act, nil
}
