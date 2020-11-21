package tracker

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// DirectorId director unique identifier
//	@ValueObject
type DirectorId struct {
	identifier domain.Identifier
}

func NewDirectorId(field, value string) (*DirectorId, error) {
	if field == "" {
		field = "director_id"
	}

	id, err := domain.NewIdentifier(field, value)
	if err != nil {
		return nil, err
	}

	return &DirectorId{identifier: *id}, nil
}

func (i DirectorId) Value() string {
	return i.identifier.Value()
}
