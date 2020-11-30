package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
)

// Creator creates an actor and assigns it to a User catalog
//	@ApplicationService
type Creator struct {
	repository actor.Repository
	eventBus   domain.EventBus
}

func NewCreator(repository actor.Repository, bus domain.EventBus) *Creator {
	return &Creator{
		repository: repository,
		eventBus:   bus,
	}
}

func (c Creator) Create(ctx context.Context, id actor.Id, userId actor.UserId, displayName actor.DisplayName,
	picture actor.Picture) error {
	if act, _ := newFinder(c.repository).Find(ctx, id); act != nil {
		return actor.AlreadyExists
	}

	act := actor.NewActor(id, userId, displayName, picture)
	if err := c.repository.Save(ctx, *act); err != nil {
		return err
	}

	return c.eventBus.Publish(ctx, act.PullEvents()...)
}
