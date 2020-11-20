package movie

// DisplayName film name which will be displayed to users
//	@ValueObject
type DisplayName struct {
	value string
}

func NewDisplayName(value string) (*DisplayName, error) {
	name := new(DisplayName)
	if err := name.ensureRequired(value); err != nil {
		return nil, err
	} else if err := name.ensureMaxLength(value); err != nil {
		return nil, err
	}

	return name, nil
}

func (n DisplayName) ensureRequired(value string) error {
	//	rules
	//	a.	display_name.length >= 1
	if value == "" {
		return NameRequired
	}

	return nil
}

func (n DisplayName) ensureMaxLength(value string) error {
	//	rules
	//	a.	display_name.length <= 128
	if len(value) > 128 {
		return NameAboveMaxLength
	}

	return nil
}

func (n DisplayName) Value() string {
	return n.value
}
