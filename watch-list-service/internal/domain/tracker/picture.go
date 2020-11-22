package tracker

import (
	"strings"

	"github.com/neutrinocorp/ddderr"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Picture image referenced by URL
//	@ValueObject
type Picture struct {
	url domain.Url
}

func NewPicture(field, value string) (*Picture, error) {
	picture := new(Picture)
	if err := picture.ensureValidPicture(field, value); value != "" && err != nil {
		return nil, err
	}
	url, err := domain.NewUrl("movie_picture", value)
	if err != nil {
		return nil, err
	}
	picture.url = *url
	return picture, nil
}

func (p Picture) ensureValidPicture(field, value string) error {
	//	rules
	//	a.	valid formats = [png, jpg, jpeg, webp]
	isImage := strings.HasSuffix(value, ".png") || strings.HasSuffix(value, ".jpg") ||
		strings.HasSuffix(value, ".jpeg") || strings.HasSuffix(value, ".webp")

	if !isImage {
		return ddderr.NewInvalidFormat(field, "jpg, jpeg, webp and png")
	}

	return nil
}

func (p Picture) Value() string {
	return p.url.Value()
}
