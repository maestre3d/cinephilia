package bus

import (
	"context"

	"github.com/neutrinocorp/ddderr"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

type InMemorySyncQuery struct {
	handlerMap map[string]domain.QueryHandler
}

func NewInMemorySyncQuery() *InMemorySyncQuery {
	return &InMemorySyncQuery{handlerMap: map[string]domain.QueryHandler{}}
}

func (q *InMemorySyncQuery) RegisterHandler(query domain.Query, handler domain.QueryHandler) error {
	if _, ok := q.handlerMap[query.Id()]; ok {
		return ddderr.NewAlreadyExists(nil, "query")
	}

	q.handlerMap[query.Id()] = handler
	return nil
}

func (q InMemorySyncQuery) Ask(ctx context.Context, query domain.Query) (interface{}, error) {
	if _, ok := q.handlerMap[query.Id()]; !ok {
		return nil, ddderr.NewNotFound(nil, "query")
	}

	return q.handlerMap[query.Id()].Invoke(ctx, query)
}
