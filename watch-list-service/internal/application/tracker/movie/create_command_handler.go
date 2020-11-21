package movie

import (
	"context"

	"github.com/hashicorp/go-multierror"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
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
	var resultErr *multierror.Error
	id, err := movie.NewMovieId(command.Id)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	userId, err := movie.NewUserId(command.UserId)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	categoryId, err := movie.NewCategoryId(command.CategoryId)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	directorId, err := movie.NewDirectorId(command.DirectorId)
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
	year, err := movie.NewYear(uint32(command.Year))
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
		CategoryId:  *categoryId,
		DirectorId:  *directorId,
		DisplayName: *displayName,
		Description: *description,
		Year:        *year,
		Picture:     *picture,
		WatchUrl:    *watchUrl,
		CrawlUrl:    movie.CrawlUrl{},
	})
}
