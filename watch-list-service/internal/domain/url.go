package domain

import (
	"net/url"

	"github.com/neutrinocorp/ddderr"
)

// Url external resource URL
//	@ValueObject
type Url struct {
	value string
}

func NewUrl(field, value string) (*Url, error) {
	if field == "" {
		field = "url"
	}

	u := new(Url)
	if err := u.ensureMaxLength(field, value); err != nil {
		return nil, err
	} else if err := u.ensureValidURL(field, value); value != "" && err != nil {
		return nil, err
	}
	u.value = value
	return u, nil
}

func (u Url) ensureMaxLength(field, value string) error {
	//	rules
	//	a.	value.length <= 2048
	if len(value) > 2048 {
		return ddderr.NewOutOfRange(field, "0", "2048")
	}

	return nil
}

func (u Url) ensureValidURL(field, value string) error {
	//	rules
	//	a.	valid URL using package (if value > 0)
	if _, err := url.Parse(value); err != nil {
		return ddderr.NewInvalidFormat(field, "url")
	}

	return nil
}

func (u Url) Value() string {
	return u.value
}
