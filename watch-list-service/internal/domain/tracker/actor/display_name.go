package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"

// DisplayName actor/actress name which will be displayed
//	@ValueObject
type DisplayName struct {
	name tracker.DisplayName
}

func NewDisplayName(value string) (*DisplayName, error) {
	displayName := new(DisplayName)

	if err := displayName.ensureIsRequired(value); err != nil {
		return nil, err
	}

	name, err := tracker.NewDisplayName("actor_display_name", value)
	if err != nil {
		return nil, err
	}
	displayName.name = *name
	return displayName, nil
}

func (n DisplayName) ensureIsRequired(value string) error {
	if value == "" {
		return DisplayNameRequired
	}

	return nil
}

func (n DisplayName) Value() string {
	return n.name.Value()
}
