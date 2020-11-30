package domain

// PageSize set a limit of a page when fetching data
//	@ValueObject
type PageSize struct {
	value int32
}

func NewPageSize(value int32) *PageSize {
	return &PageSize{value: value}
}

func (s PageSize) Value() int32 {
	if s.value == 0 {
		// return default value if nil/empty
		return 10
	}

	return s.value
}
