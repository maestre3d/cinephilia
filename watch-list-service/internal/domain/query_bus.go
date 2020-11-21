package domain

import "context"

// QueryBus receives queries and dispatch them to their respective query handlers
//	@Port
//	@Query
//	@Bus
type QueryBus interface {
	RegisterHandler(query Query, handler QueryHandler) error
	Ask(ctx context.Context, query Query) (interface{}, error)
}
