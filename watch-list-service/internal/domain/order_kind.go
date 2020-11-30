package domain

const (
	asc  = "asc"
	desc = "desc"
	none = "none"
)

// OrderKind criteria ordering kind
//	@ValueObject
type OrderKind struct {
	value string
}

func NewOrderKind(value string) *OrderKind {
	return &OrderKind{value: value}
}

func (t OrderKind) IsNone() bool {
	return t.value == none
}

func (t OrderKind) IsAsc() bool {
	return t.value == asc
}

func (t OrderKind) IsDesc() bool {
	return t.value == desc
}

func (t OrderKind) Value() string {
	return t.value
}
