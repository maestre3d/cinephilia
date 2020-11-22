package director

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Director film director
//	@Entity
//	@AggregateRoot
type Director struct {
	id          Id
	userId      UserId
	displayName DisplayName
	picture     Picture

	// extends aggregate root
	events []domain.Event
}

// TODO: Add application services

func NewDirector(id Id, userId UserId, displayName DisplayName, picture Picture) *Director {
	dir := &Director{
		id:          id,
		userId:      userId,
		displayName: displayName,
		picture:     picture,
		events:      []domain.Event{},
	}
	dir.record(NewDirectorCreated(id.Value(), userId.Value(), displayName.Value(), picture.Value()))
	return dir
}

func (d *Director) record(event ...domain.Event) {
	d.events = append(d.events, event...)
}

func (d *Director) PullEvents() []domain.Event {
	memoized := d.events
	d.events = []domain.Event{}

	return memoized
}
