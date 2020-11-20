package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

// Creator creates a movie using basic data
//	@ApplicationService
type Creator struct {
	repository movie.Repository
}

func NewCreator(repository movie.Repository) *Creator {
	return &Creator{repository: repository}
}

func (c Creator) Create(ctx context.Context, id movie.Id, name movie.DisplayName, description movie.Description,
	userId movie.UserId, categoryId movie.CategoryId) error {
	if mov, _ := NewFinder(c.repository).Find(ctx, id); mov != nil {
		return movie.AlreadyExists
	}

	return c.repository.Save(ctx, *movie.NewMovie(id, name, description, userId, categoryId))
}
