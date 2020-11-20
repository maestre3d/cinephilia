package domain

import "fmt"

// Filter filter aggregates by specific fields
//	@ValueObject
type Filter struct {
	field    string
	operator FilterOperator
	value    string
}

func NewFilter(field, value string, operator FilterOperator) *Filter {
	return &Filter{
		field:    field,
		operator: operator,
		value:    value,
	}
}

func NewFilterFromPrimitives(field, operator, value string) *Filter {
	return &Filter{
		field:    field,
		operator: *NewFilterOperatorFromPrimitive(operator),
		value:    value,
	}
}

func (f Filter) Serialize() string {
	return fmt.Sprintf("%s.%s.%s", f.field, f.operator, f.value)
}

// Filters list of Filter
//	@ValueObject
type Filters struct {
	values []Filter
}

func NewFilters(filter ...Filter) *Filters {
	return &Filters{values: filter}
}

func NewFiltersFromPrimitives(filters []map[string]string) *Filters {
	filterSlice := make([]Filter, len(filters))
	for _, filter := range filters {
		filterSlice = append(filterSlice,
			*NewFilterFromPrimitives(filter["field"], filter["operator"], filter["value"]))
	}

	return &Filters{
		values: filterSlice,
	}
}

func (f Filters) Values() []Filter {
	return f.values
}
