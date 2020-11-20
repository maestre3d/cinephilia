package tracker

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// CategoryId category unique identifier
//	@ValueObject
type CategoryId struct {
	identifier domain.Identifier
}

func NewCategoryId(field, value string) (*CategoryId, error) {
	if field == "" {
		field = "category_id"
	}

	id, err := domain.NewIdentifier(field, value)
	if err != nil {
		return nil, err
	}
	return &CategoryId{identifier: *id}, nil
}

func (i CategoryId) Value() string {
	return i.identifier.Value()
}
