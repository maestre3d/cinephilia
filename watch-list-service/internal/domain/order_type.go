package domain

const (
	asc  = "asc"
	desc = "desc"
	none = "none"
)

// OrderType criteria ordering kind
//	@ValueObject
type OrderType struct {
	value string
}

func NewOrderType(value string) *OrderType {
	return &OrderType{value: value}
}

func (t OrderType) IsNone() bool {
	return t.value == none
}

func (t OrderType) IsAsc() bool {
	return t.value == asc
}

func (t OrderType) IsDesc() bool {
	return t.value == desc
}

func (t OrderType) Value() string {
	return t.value
}
