package movie

// Description detailed text about a film
//	@ValueObject
type Description struct {
	value string
}

func NewDescription(value string) (*Description, error) {
	description := new(Description)
	if err := description.ensureMaxLength(value); err != nil {
		return nil, err
	}
	description.value = value

	return description, nil
}

func (d Description) ensureMaxLength(value string) error {
	//	rules
	//	a.	description.length <= 512
	if len(value) > 512 {
		return DescriptionAboveMaxLength
	}

	return nil
}

func (d Description) Value() string {
	return d.value
}

func (d Description) NullableValue() *string {
	if d.Value() == "" {
		return nil
	}

	val := d.value
	return &val
}
