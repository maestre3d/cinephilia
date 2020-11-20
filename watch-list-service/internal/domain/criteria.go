package domain

type Filters []Filter

// Criteria search aggregates by an order, limit, page_token and specific fields
//	@DTO
type Criteria struct {
	limit     uint16
	pageToken string
	order     Order
	filters   Filters
}
