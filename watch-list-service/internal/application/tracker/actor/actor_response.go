package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"

// Response actor search response
//	@DTO
type Response struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	DisplayName string  `json:"display_name"`
	Picture     *string `json:"picture"`
}

func newResponse(act actor.Actor) *Response {
	return &Response{
		Id:          act.Id().Value(),
		UserId:      act.User().Value(),
		DisplayName: act.Name().Value(),
		Picture:     act.Picture().NullableValue(),
	}
}
