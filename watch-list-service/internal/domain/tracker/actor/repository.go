package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, actor Actor) error
	Search(ctx context.Context, id Id) (*Actor, error)
	SearchByCriteria(ctx context.Context, criteria domain.Criteria) ([]*Actor, domain.PageToken, error)
}
