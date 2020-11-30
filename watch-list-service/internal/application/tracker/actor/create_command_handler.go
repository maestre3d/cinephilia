package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
	"github.com/neutrinocorp/ddderr"
)

// CreateCommandHandler handles actor creations
//	@Command
type CreateCommandHandler struct {
	creator *Creator
}

func NewCreateCommandHandler(creator *Creator) *CreateCommandHandler {
	return &CreateCommandHandler{creator: creator}
}

func (h CreateCommandHandler) Invoke(ctx context.Context, command domain.Command) error {
	cmd, ok := command.(CreateCommand)
	if !ok {
		return ddderr.NewInvalidFormat("command", command.Id())
	}

	id, err := actor.NewId(cmd.ActorId)
	if err != nil {
		return err
	}
	userId, err := actor.NewUserId(cmd.UserId)
	if err != nil {
		return err
	}
	displayName, err := actor.NewDisplayName(cmd.DisplayName)
	if err != nil {
		return err
	}
	picture, err := actor.NewPicture(cmd.Picture)
	if err != nil {
		return err
	}

	return h.creator.Create(ctx, *id, *userId, *displayName, *picture)
}
