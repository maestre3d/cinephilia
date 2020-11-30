package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
)

// Updater updates the given actor
//	@ApplicationService
type Updater struct {
	repository actor.Repository
	eventBus   domain.EventBus
}

func NewUpdater(repository actor.Repository, bus domain.EventBus) *Updater {
	return &Updater{
		repository: repository,
		eventBus:   bus,
	}
}

func (u Updater) Update(ctx context.Context, id actor.Id, displayName actor.DisplayName,
	picture actor.Picture) error {
	act, err := newFinder(u.repository).Find(ctx, id)
	if err != nil {
		// infra error or not found thrown/return
		return err
	}

	switch {
	case displayName.Value() != "":
		act.ChangeName(displayName)
		fallthrough
	case picture.Value() != "":
		act.ChangePicture(picture)
	}

	act.Update()
	if err := u.repository.Save(ctx, *act); err != nil {
		return err
	}

	return u.eventBus.Publish(ctx, act.PullEvents()...)
}
