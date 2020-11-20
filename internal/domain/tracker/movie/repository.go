package movie

import "context"

// Criteria search movies by order, limit, offset and specific fields
//	@DTO
type Criteria struct{}

// Repository interacts with the respective persistence infrastructure for movie aggregate
//	@Port
type Repository interface {
	Save(ctx context.Context, movie Movie) error
	Search(ctx context.Context, movieId Id) (*Movie, error)
	SearchByCriteria(ctx context.Context, criteria Criteria) ([]*Movie, error)
}
