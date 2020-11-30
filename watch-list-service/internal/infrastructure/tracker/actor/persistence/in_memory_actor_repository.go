package persistence

import (
	"sync"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
)

type InMemoryActorRepository struct {
	db map[string]*actor.Actor
	mu *sync.RWMutex
}

func NewInMemoryActorRepository() *InMemoryActorRepository {
	return &InMemoryActorRepository{
		db: map[string]*actor.Actor{},
		mu: new(sync.RWMutex),
	}
}
