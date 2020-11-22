package marshaler

import "time"

//	@DTO
type Director struct {
	Id          string
	UserId      string
	DisplayName string
	Picture     *string
	CreateTime  time.Time
	UpdateTime  time.Time
	Active      bool
}

func (d Director) SafePicture() string {
	if d.Picture == nil {
		return ""
	}

	return *d.Picture
}
