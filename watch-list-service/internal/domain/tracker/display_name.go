package tracker

import "github.com/neutrinocorp/ddderr"

// DisplayName name which will be displayed to users
//	@ValueObject
type DisplayName struct {
	value string
}

func NewDisplayName(field, value string) (*DisplayName, error) {
	name := new(DisplayName)
	if err := name.ensureMaxLength(field, value); err != nil {
		return nil, err
	}
	name.value = value

	return name, nil
}

func (n DisplayName) ensureMaxLength(field, value string) error {
	//	rules
	//	a.	display_name.length <= 128
	if len(value) > 128 {
		return ddderr.NewOutOfRange(field, "1", "128")
	}

	return nil
}

func (n DisplayName) Value() string {
	return n.value
}
