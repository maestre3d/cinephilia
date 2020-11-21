package domain

import "context"

// QueryHandler receives Query(s) and executes the fetching operation
//	@Port
type QueryHandler interface {
	Invoke(ctx context.Context, query Query) (interface{}, error)
}
