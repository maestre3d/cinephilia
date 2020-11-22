package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// MovieId film unique identifier
//	@ValueObject
type Id struct {
	// extends Identifier
	value domain.Identifier
}

func NewMovieId(value string) (*Id, error) {
	movieId := new(Id)
	if err := movieId.ensureIsRequired(value); err != nil {
		return nil, err
	}

	id, err := domain.NewIdentifier("movie_id", value)
	if err != nil {
		return nil, err
	}

	movieId.value = *id
	return movieId, nil
}

func (i Id) ensureIsRequired(value string) error {
	//	rules
	//	a.	value > 1
	if value == "" {
		return IdRequired
	}

	return nil
}

func (i Id) Value() string {
	return i.value.Value()
}
