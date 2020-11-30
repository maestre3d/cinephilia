package actor

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Id actor unique identifier
//	@ValueObject
type Id struct {
	identifier domain.Identifier
}

func NewId(value string) (*Id, error) {
	Id := new(Id)

	if err := Id.ensureIsRequired(value); err != nil {
		return nil, err
	}

	id, err := domain.NewIdentifier("actor_id", value)
	if err != nil {
		return nil, err
	}
	Id.identifier = *id

	return Id, nil
}

func (i Id) ensureIsRequired(value string) error {
	if value == "" {
		return IdRequired
	}

	return nil
}

func (i Id) Value() string {
	return i.identifier.Value()
}
