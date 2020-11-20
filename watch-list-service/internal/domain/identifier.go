package domain

import "github.com/neutrinocorp/ddderr"

// Identifier high-cardinal field to identify entities/aggregates
//	@ValueObject
type Identifier struct {
	value string
}

func NewIdentifier(field, value string) (*Identifier, error) {
	if field == "" {
		field = "id"
	}
	id := new(Identifier)
	if err := id.ensureValidNanoId(field, value); err != nil {
		return nil, err
	}

	id.value = value
	return id, nil
}

func (i Identifier) ensureValidNanoId(field, value string) error {
	//	rules
	//	a.	16 <= value.length >= 128
	if len(value) < 16 || len(value) > 128 {
		return ddderr.NewInvalidFormat(field, "nano id")
	}

	return nil
}

func (i Identifier) Value() string {
	return i.value
}
