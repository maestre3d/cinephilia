package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// UpdateCommand request an Actor update
//	@Command
//	@DTO
type UpdateCommand struct {
	ActorId     string
	DisplayName string
	Picture     string
}

func (c UpdateCommand) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.actor.update"
}
