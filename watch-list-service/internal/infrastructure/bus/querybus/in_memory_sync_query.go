package querybus

import (
	"context"
	"errors"
	"sync"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

type InMemorySyncQuery struct {
	mu         sync.Mutex
	handlerMap map[string]domain.QueryHandler
}

func NewInMemorySyncQuery() *InMemorySyncQuery {
	return &InMemorySyncQuery{handlerMap: map[string]domain.QueryHandler{}}
}

func (q *InMemorySyncQuery) RegisterHandler(query domain.Query, handler domain.QueryHandler) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if _, ok := q.handlerMap[query.Id()]; ok {
		return errors.New("query already exists")
	}

	q.handlerMap[query.Id()] = handler
	return nil
}

func (q *InMemorySyncQuery) Ask(ctx context.Context, query domain.Query) (interface{}, error) {
	if _, ok := q.handlerMap[query.Id()]; !ok {
		return nil, errors.New("query does not exists")
	}

	return q.handlerMap[query.Id()].Invoke(ctx, query)
}
