package movie

// Year year a film was released
//	@ValueObject
type Year struct {
	value uint32
}

func NewYear(value uint32) (*Year, error) {
	year := new(Year)
	if err := year.ensureMaxLength(value); err != nil {
		return nil, err
	} else if err := year.ensureMinLength(value); err != nil {
		return nil, err
	}
	year.value = value

	return year, nil
}

func (y Year) ensureMaxLength(value uint32) error {
	//	rules
	//	a.	0 >= year.length <= 3000
	if value > 3000 {
		return YearAboveMaxLength
	}

	return nil
}

func (y Year) ensureMinLength(value uint32) error {
	//	rules
	//	a.	0 >= year.length <= 3000
	if value < 0 {
		return YearBelowMinLength
	}

	return nil
}

func (y Year) Value() uint32 {
	return y.value
}

func (y Year) NullableValue() *uint32 {
	if y.Value() == 0 {
		return nil
	}
	return &y.value
}
