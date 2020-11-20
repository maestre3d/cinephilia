package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/internal/domain/tracker/movie"
)

// Finder search for a Movie
//	@DomainService
type Finder struct {
	repository movie.Repository
}

func NewFinder(repository movie.Repository) *Finder {
	return &Finder{repository: repository}
}

func (f Finder) Find(ctx context.Context, movieId movie.Id) (*movie.Movie, error) {
	return f.repository.Search(ctx, movieId)
}
