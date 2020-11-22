package director

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"
)

// Picture director portrait image referenced by URL
//	@ValueObject
type Picture struct {
	value tracker.Picture
}

func NewPicture(value string) (*Picture, error) {
	picture := new(Picture)
	pic, err := tracker.NewPicture("director_picture", value)
	if err != nil {
		return nil, err
	}
	picture.value = *pic
	return picture, nil
}

func (p Picture) IsEmpty() bool {
	return p.Value() == ""
}

func (p Picture) Value() string {
	return p.value.Value()
}

func (p Picture) NullableValue() *string {
	if p.Value() == "" {
		return nil
	}
	val := p.Value()
	return &val
}
