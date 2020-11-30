package movie

import (
	"context"
	"strconv"

	"github.com/hashicorp/go-multierror"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/neutrinocorp/ddderr"
)

// CreateCommandHandler receives CreateCommand(s) and creates a Movie
//	@CommandHandler
type CreateCommandHandler struct {
	creator *Creator
}

func NewCreateCommandHandler(creator *Creator) *CreateCommandHandler {
	return &CreateCommandHandler{creator: creator}
}

func (h CreateCommandHandler) Invoke(ctx context.Context, cmd domain.Command) error {
	command, ok := cmd.(CreateCommand)
	if !ok {
		return ddderr.NewInvalidFormat("command", cmd.Id())
	}

	var resultErr *multierror.Error
	id, err := movie.NewMovieId(command.MovieId)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	userId, err := movie.NewUserId(command.UserId)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	displayName, err := movie.NewDisplayName(command.DisplayName)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	description, err := movie.NewDescription(command.Description)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	yearInt, err := strconv.ParseInt(command.Year, 10, 32)
	if err != nil {
		return ddderr.NewInvalidFormat("year", "integer")
	}
	year, err := movie.NewYear(uint32(yearInt))
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	picture, err := movie.NewPicture(command.Picture)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	watchUrl, err := movie.NewWatchUrl(command.WatchUrl)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}

	if resultErr != nil && resultErr.ErrorOrNil() != nil {
		return resultErr
	}
	return h.creator.Create(ctx, CreatorArgs{
		Id:          *id,
		UserId:      *userId,
		DisplayName: *displayName,
		Description: *description,
		Year:        *year,
		Picture:     *picture,
		WatchUrl:    *watchUrl,
		CrawlUrl:    movie.CrawlUrl{},
	})
}
