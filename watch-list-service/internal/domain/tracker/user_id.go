package tracker

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// UserId user unique identifier
//	@ValueObject
type UserId struct {
	value domain.Identifier
}

func NewUserId(field, value string) (*UserId, error) {
	if field == "" {
		field = "user_id"
	}

	id, err := domain.NewIdentifier(field, value)
	if err != nil {
		return nil, err
	}
	return &UserId{value: *id}, nil
}

func (i UserId) Value() string {
	return i.value.Value()
}
