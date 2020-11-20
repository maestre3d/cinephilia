package domain

// Criteria search aggregates by an order, limit, page_token and specific fields
//	@DTO
//	@DSL (Domain Selection Language)
type Criteria struct {
	limit     uint16
	pageToken string
	order     Order
	filters   Filters
}

func NewCriteria(limit uint16, pageToken string, order Order, filters Filters) *Criteria {
	return &Criteria{
		limit:     limit,
		pageToken: pageToken,
		order:     order,
		filters:   filters,
	}
}

func (c Criteria) Limit() uint16 {
	return c.limit
}

func (c Criteria) PageToken() string {
	return c.pageToken
}

func (c Criteria) Order() Order {
	return c.order
}

func (c Criteria) Filters() Filters {
	return c.filters
}

func (c Criteria) HasFilters() bool {
	return len(c.filters.Values()) > 1
}
