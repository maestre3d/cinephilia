package director

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"

// DisplayName director name which will be displayed to users
//	@ValueObject
type DisplayName struct {
	value tracker.DisplayName
}

func NewDisplayName(value string) (*DisplayName, error) {
	name := new(DisplayName)
	if err := name.ensureRequired(value); err != nil {
		return nil, err
	}
	displayName, err := tracker.NewDisplayName("director_display_name", value)
	if err != nil {
		return nil, err
	}

	name.value = *displayName
	return name, nil
}

func (n DisplayName) ensureRequired(value string) error {
	//	rules
	//	a.	display_name.length >= 1
	if value == "" {
		return DisplayNameRequired
	}

	return nil
}

func (n DisplayName) Value() string {
	return n.value.Value()
}
