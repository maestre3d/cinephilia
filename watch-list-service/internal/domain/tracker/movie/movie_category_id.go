package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"

// CategoryId category unique identifier, assigns a movie to a category
//	@ValueObject
type CategoryId struct {
	// extends CategoryId
	value tracker.CategoryId
}

func NewCategoryId(value string) (*CategoryId, error) {
	id, err := tracker.NewCategoryId("movie_category_id", value)
	if err != nil {
		return nil, err
	}

	return &CategoryId{value: *id}, nil
}

func (i CategoryId) IsEmpty() bool {
	return i.value.Value() == ""
}

func (i CategoryId) Value() string {
	return i.value.Value()
}

func (i CategoryId) NullableValue() *string {
	if i.Value() == "" {
		return nil
	}
	val := i.Value()
	return &val
}
