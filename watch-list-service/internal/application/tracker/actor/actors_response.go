package actor

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/actor"
)

// ActorsResponse list of actors
//	@DTO
type ActorsResponse struct {
	Actors        []*Response `json:"actors"`
	NextPageToken string      `json:"next_page_token"`
	Total         int         `json:"total"`
}

func newActorsResponse(nextPage domain.PageToken, acts ...*actor.Actor) *ActorsResponse {
	res := &ActorsResponse{Actors: make([]*Response, 0), NextPageToken: nextPage.Value(), Total: len(acts)}
	for _, act := range acts {
		res.Actors = append(res.Actors, newResponse(*act))
	}

	return res
}
