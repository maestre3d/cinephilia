package movie

import (
	"strings"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Picture film cover/poster image referenced by URL
//	@ValueObject
type Picture struct {
	url domain.Url
}

func NewPicture(value string) (*Picture, error) {
	picture := new(Picture)
	if err := picture.ensureValidPicture(value); value != "" && err != nil {
		return nil, err
	}
	url, err := domain.NewUrl("movie_picture", value)
	if err != nil {
		return nil, err
	}
	picture.url = *url
	return picture, nil
}

func (p Picture) ensureValidPicture(value string) error {
	//	rules
	//	a.	valid formats = [png, jpg, jpeg, webp]
	isImage := strings.HasSuffix(value, ".png") || strings.HasSuffix(value, ".jpg") ||
		strings.HasSuffix(value, ".jpeg") || strings.HasSuffix(value, ".webp")

	if !isImage {
		return InvalidPictureExtension
	}

	return nil
}

func (p Picture) IsEmpty() bool {
	return p.Value() == ""
}

func (p Picture) Value() string {
	return p.url.Value()
}

func (p Picture) NullableValue() *string {
	if p.Value() == "" {
		return nil
	}

	val := p.Value()
	return &val
}
