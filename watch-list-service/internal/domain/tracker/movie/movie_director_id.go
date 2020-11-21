package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"

// DirectorId film director unique identifier
//	@ValueObject
type DirectorId struct {
	identifier tracker.DirectorId
}

func NewDirectorId(value string) (*DirectorId, error) {
	id, err := tracker.NewDirectorId("movie_director_id", value)
	if err != nil {
		return nil, err
	}

	return &DirectorId{identifier: *id}, nil
}

func (i DirectorId) IsEmpty() bool {
	return i.Value() == ""
}

func (i DirectorId) Value() string {
	return i.identifier.Value()
}

func (i DirectorId) NullableValue() *string {
	if i.Value() == "" {
		return nil
	}
	val := i.Value()
	return &val
}
