package domain

// PageToken set an start point to fetch data
//	@ValueObject
type PageToken struct {
	value string
}

func NewPageToken(value string) *PageToken {
	return &PageToken{value: value}
}

func (t PageToken) Value() string {
	return t.value
}
