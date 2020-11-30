package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// Actor man or woman who appears in films
//	@AggregateRoot
//	@Entity
type Actor struct {
	id          Id
	user        UserId
	displayName DisplayName
	picture     Picture

	events []domain.Event
}

func NewActor(id Id, userId UserId, displayName DisplayName, picture Picture) *Actor {
	actor := &Actor{
		id:          id,
		user:        userId,
		displayName: displayName,
		picture:     picture,
		events:      make([]domain.Event, 0),
	}
	actor.record(newActorCreated(actor.id.Value(), actor.user.Value(), actor.displayName.Value(),
		actor.Picture().Value()))
	return actor
}

func (a *Actor) Update() {
	a.record(newActorUpdated(a.id.Value(), a.user.Value(), a.displayName.Value(),
		a.Picture().Value()))
}

func (a *Actor) ChangeName(displayName DisplayName) {
	a.displayName = displayName
}

func (a *Actor) ChangePicture(picture Picture) {
	a.picture = picture
}

func (a *Actor) record(event ...domain.Event) {
	a.events = append(a.events, event...)
}

func (a *Actor) PullEvents() []domain.Event {
	memoized := a.events
	a.events = []domain.Event{}

	return memoized
}

func (a Actor) Id() Id {
	return a.id
}

func (a Actor) User() UserId {
	return a.user
}

func (a Actor) Name() DisplayName {
	return a.displayName
}

func (a Actor) Picture() Picture {
	return a.picture
}
