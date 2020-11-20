package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

// Finder search for a Movie
//	@DomainService
type Finder struct {
	repository movie.Repository
}

func NewFinder(repository movie.Repository) *Finder {
	return &Finder{repository: repository}
}

func (f Finder) Find(ctx context.Context, movieId movie.Id) (*MovieResponse, error) {
	mov, err := f.repository.Search(ctx, movieId)
	if err != nil {
		return nil, err
	} else if mov == nil {
		return nil, movie.NotExists
	}

	return newMovieResponseFromAggregate(*mov), nil
}
