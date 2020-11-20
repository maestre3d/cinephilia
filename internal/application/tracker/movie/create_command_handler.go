package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/internal/domain/tracker/movie"
)

// CreateCommandHandler receives CreateCommand(s) and creates a Movie
//	@CommandHandler
type CreateCommandHandler struct {
	creator *Creator
}

func NewCreateCommandHandler(creator *Creator) *CreateCommandHandler {
	return &CreateCommandHandler{creator: creator}
}

func (h CreateCommandHandler) Invoke(ctx context.Context, command CreateCommand) error {
	id, err := movie.NewMovieId(command.Id)
	if err != nil {
		return err
	}
	displayName, err := movie.NewDisplayName(command.DisplayName)
	if err != nil {
		return err
	}
	description, err := movie.NewDescription(command.Description)
	if err != nil {
		return err
	}
	userId, err := movie.NewUserId(command.UserId)
	if err != nil {
		return err
	}

	return h.creator.Create(ctx, *id, *displayName, *description, *userId)
}
