package director

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"
)

// Id director unique identifier
//	@ValueObject
type Id struct {
	identifier tracker.DirectorId
}

func NewDirectorId(value string) (*Id, error) {
	directorId := new(Id)
	if err := directorId.ensureIsRequired(value); err != nil {
		return nil, err
	}

	id, err := tracker.NewDirectorId("", value)
	if err != nil {
		return nil, err
	}

	directorId.identifier = *id
	return directorId, nil
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
	return i.identifier.Value()
}
