package actor

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/neutrinocorp/ddderr"
)

// SearchQueryHandler retrieves a list of actors
//	@QueryHandler
type SearchQueryHandler struct {
	searcher *Searcher
}

func NewSearchQueryHandler(searcher *Searcher) *SearchQueryHandler {
	return &SearchQueryHandler{searcher: searcher}
}

func (h SearchQueryHandler) Invoke(ctx context.Context, query domain.Query) (interface{}, error) {
	q, ok := query.(SearchQuery)
	if !ok {
		return nil, ddderr.NewInvalidFormat("query", query.Id())
	}

	domain.NewOrder(domain.OrderBy(q.OrderBy), *domain.NewOrderKind(q.OrderKind))
	criteria := domain.NewCriteria(*domain.NewPageSize(q.PageSize), *domain.NewPageToken(q.PageToken),
		*domain.NewOrder(domain.OrderBy(q.OrderBy), *domain.NewOrderKind(q.OrderKind)),
		*domain.NewFiltersFromPrimitives(q.Filters))
	return h.searcher.Search(ctx, *criteria)
}
