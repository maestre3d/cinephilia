package actor

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// FindQuery search an actor by unique identifier
//	@DTO
//	@Query
type FindQuery struct {
	ActorId string
}

func (q FindQuery) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.actor.find"
}
