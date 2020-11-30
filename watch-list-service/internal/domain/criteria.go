package domain

// Criteria search aggregates by an order, limit, page_token and specific fields
//	@DTO
//	@DSL (Domain Selection Language)
type Criteria struct {
	pageSize  PageSize
	pageToken PageToken
	order     Order
	filters   Filters
}

func NewCriteria(pageSize PageSize, pageToken PageToken, order Order, filters Filters) *Criteria {
	return &Criteria{
		pageSize:  pageSize,
		pageToken: pageToken,
		order:     order,
		filters:   filters,
	}
}

func (c Criteria) PageSize() PageSize {
	return c.pageSize
}

func (c Criteria) PageToken() PageToken {
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

func (c Criteria) HasPageToken() bool {
	return c.pageToken.Value() != ""
}
