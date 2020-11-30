package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
)

// Searcher fetch an actor by criteria
//	@ApplicationService
type Searcher struct {
	repository actor.Repository
}

func NewSearcher(repository actor.Repository) *Searcher {
	return &Searcher{repository: repository}
}

func (s Searcher) Search(ctx context.Context, criteria domain.Criteria) (*ActorsResponse, error) {
	acts, next, err := s.repository.SearchByCriteria(ctx, criteria)
	if err != nil {
		return nil, err
	} else if len(acts) == 0 {
		return nil, actor.NotFound
	}

	return newActorsResponse(next, acts...), nil
}
