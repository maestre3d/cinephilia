package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
	"github.com/neutrinocorp/ddderr"
)

// UpdateCommandHandler executes actor updates
//	@CommandHandler
type UpdateCommandHandler struct {
	updater *Updater
}

func NewUpdateCommandHandler(updater *Updater) *UpdateCommandHandler {
	return &UpdateCommandHandler{updater: updater}
}

func (h UpdateCommandHandler) Invoke(ctx context.Context, command domain.Command) error {
	cmd, ok := command.(UpdateCommand)
	if !ok {
		return ddderr.NewInvalidFormat("command", command.Id())
	}

	id, err := actor.NewId(cmd.ActorId)
	if err != nil {
		return err
	}
	// set default value
	displayName := new(actor.DisplayName)
	if cmd.DisplayName != "" {
		displayName, err = actor.NewDisplayName(cmd.DisplayName)
		if err != nil {
			return err
		}
	}
	picture, err := actor.NewPicture(cmd.Picture)
	if err != nil {
		return err
	}

	return h.updater.Update(ctx, *id, *displayName, *picture)
}
