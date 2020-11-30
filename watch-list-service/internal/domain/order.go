package domain

import "fmt"

// OrderBy criteria ordering field
//	@ValueObject
type OrderBy string

// Order set ordering to an aggregate search by criteria
//	@ValueObject
type Order struct {
	by   OrderBy
	kind OrderKind
}

func NewOrder(orderBy OrderBy, orderKind OrderKind) *Order {
	return &Order{
		by:   orderBy,
		kind: orderKind,
	}
}

func (o Order) By() OrderBy {
	return o.by
}

func (o Order) Kind() OrderKind {
	return o.kind
}

func (o Order) HasOrder() bool {
	return !o.kind.IsNone()
}

func (o Order) Serialize() string {
	return fmt.Sprintf("%s.%s", o.by, o.kind.Value())
}
