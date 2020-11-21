package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// WatchUrl film external URL for quick user access
//	@ValueObject
type WatchUrl struct {
	// extend domain.Url
	url domain.Url
}

func NewWatchUrl(value string) (*WatchUrl, error) {
	url, err := domain.NewUrl("watch_url", value)
	if err != nil {
		return nil, err
	}

	return &WatchUrl{url: *url}, nil
}

func (u WatchUrl) IsEmpty() bool {
	return u.Value() == ""
}

func (u WatchUrl) Value() string {
	return u.url.Value()
}

func (u WatchUrl) NullableValue() *string {
	if u.Value() == "" {
		return nil
	}
	val := u.Value()
	return &val
}
