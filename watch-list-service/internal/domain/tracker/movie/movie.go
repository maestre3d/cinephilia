package movie

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Movie is a film a user either wants to see or has already seen
//	@AggregateRoot
//	@Entity
type Movie struct {
	id          Id
	displayName DisplayName
	description Description
	user        UserId

	// extend aggregate root
	events []domain.Event
}

func NewMovie(id Id, name DisplayName, description Description, userId UserId) *Movie {
	mov := &Movie{
		id:          id,
		displayName: name,
		description: description,
		user:        userId,
		events:      make([]domain.Event, 0),
	}
	mov.record(NewMovieCreated(id.Value(), name.Value(), description.Value()))
	return mov
}

func (m *Movie) Create(id Id, name DisplayName, description Description, user UserId) {

}

func (m *Movie) record(event ...domain.Event) {
	m.events = append(m.events, event...)
}

func (m Movie) PullEvents() []domain.Event {
	memoizedEvents := m.events
	m.events = []domain.Event{}

	return memoizedEvents
}

func (m Movie) Id() Id {
	return m.id
}

func (m Movie) Name() DisplayName {
	return m.displayName
}

func (m Movie) Description() Description {
	return m.description
}

func (m Movie) User() UserId {
	return m.user
}
