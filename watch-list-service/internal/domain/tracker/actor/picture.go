package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"

// Picture image of an actor/actress
//	@ValueObject
type Picture struct {
	picture tracker.Picture
}

func NewPicture(value string) (*Picture, error) {
	pic, err := tracker.NewPicture("actor_picture", value)
	if err != nil {
		return nil, err
	}

	return &Picture{picture: *pic}, nil
}

func (p Picture) Value() string {
	return p.picture.Value()
}

func (p Picture) NullableValue() *string {
	if p.Value() == "" {
		return nil
	}
	val := p.Value()
	return &val
}
