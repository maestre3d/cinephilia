package movie

import "github.com/maestre3d/cinephilia/internal/domain"

// MovieId film unique identifier
//	@ValueObject
type Id struct {
	// extends Identifier
	value domain.Identifier
}

func NewMovieId(value string) (*Id, error) {
	id, err := domain.NewIdentifier("movie_id", value)
	if err != nil {
		return nil, err
	}

	return &Id{value: *id}, nil
}

func (i Id) Value() string {
	return i.value.Value()
}
