package persistence

import (
	"context"
	"sync"

	"github.com/maestre3d/cinephilia/internal/domain/tracker/movie"
)

type InMemoryMovieRepository struct {
	db map[string]*movie.Movie
	mu *sync.RWMutex
}

func NewInMemoryMovieRepository() *InMemoryMovieRepository {
	return &InMemoryMovieRepository{
		db: map[string]*movie.Movie{},
		mu: new(sync.RWMutex),
	}
}

func (i *InMemoryMovieRepository) Save(_ context.Context, movie movie.Movie) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.db[movie.Id().Value()] = &movie
	return nil
}

func (i InMemoryMovieRepository) Search(_ context.Context, movieId movie.Id) (*movie.Movie, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.db[movieId.Value()], nil
}

func (i InMemoryMovieRepository) SearchByCriteria(_ context.Context, criteria movie.Criteria) ([]*movie.Movie, error) {
	panic("implement me")
}
