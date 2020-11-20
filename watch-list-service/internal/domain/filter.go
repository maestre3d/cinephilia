package domain

// Filter filter aggregates by specific fields
//	@ValueObject
type Filter struct {
	field    string
	operator string
	value    string
}
