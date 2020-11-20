package domain

const (
	equal       = "EQUAL"
	notEqual    = "NOT_EQUAL"
	greater     = ">"
	less        = "<"
	contains    = "CONTAINS"
	notContains = "NOT_CONTAINS"
)

// FilterOperator boolean operator to a filter
//	@ValueObject
type FilterOperator struct {
	value string
}

func NewFilterOperator(value string) *FilterOperator {
	return &FilterOperator{value: value}
}

func NewFilterOperatorFromPrimitive(value string) *FilterOperator {
	switch value {
	case "=":
		return &FilterOperator{value: equal}
	case "!=":
		return &FilterOperator{notEqual}
	case ">":
		return &FilterOperator{greater}
	case "<":
		return &FilterOperator{less}
	case "CONTAINS":
		return &FilterOperator{contains}
	case "NOT_CONTAINS":
		return &FilterOperator{value: notContains}
	default:
		return nil
	}
}

func (o FilterOperator) Value() string {
	return o.value
}
