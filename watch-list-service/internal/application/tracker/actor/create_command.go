package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// CreateCommand requests an actor creation
//	@DTO
//	@Command
type CreateCommand struct {
	ActorId     string
	UserId      string
	DisplayName string
	Picture     string
}

func (c CreateCommand) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.actor.create"
}
