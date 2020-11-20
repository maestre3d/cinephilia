package domain

// OrderBy criteria ordering field
//	@ValueObject
type OrderBy string

// Order set ordering to an aggregate search by criteria
//	@ValueObject
type Order struct {
	orderBy   OrderBy
	orderType OrderType
}

func NewOrder(orderBy OrderBy, orderType OrderType) *Order {
	return &Order{
		orderBy:   orderBy,
		orderType: orderType,
	}
}

func (o Order) By() OrderBy {
	return o.orderBy
}

func (o Order) Type() OrderType {
	return o.orderType
}
