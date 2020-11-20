package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Repository interacts with the respective persistence infrastructure for movie aggregate
//	@Port
type Repository interface {
	Save(ctx context.Context, movie Movie) error
	Search(ctx context.Context, movieId Id) (*Movie, error)
	SearchByCriteria(ctx context.Context, criteria domain.Criteria) ([]*Movie, error)
}
