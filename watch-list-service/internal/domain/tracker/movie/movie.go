package movie

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// Movie is a film a user either wants to see or has already seen
//	@AggregateRoot
//	@Entity
type Movie struct {
	id          Id
	user        UserId
	categoryId  CategoryId
	director    DirectorId
	displayName DisplayName
	description Description
	year        Year

	picture  Picture
	watchUrl WatchUrl
	crawlUrl CrawlUrl

	// extend aggregate root
	events []domain.Event
}

func (m *Movie) record(event ...domain.Event) {
	m.events = append(m.events, event...)
}

func (m *Movie) PullEvents() []domain.Event {
	memoizedEvents := m.events
	m.events = []domain.Event{}

	return memoizedEvents
}

func (m Movie) Id() Id {
	return m.id
}

func (m Movie) User() UserId {
	return m.user
}

func (m Movie) Category() CategoryId {
	return m.categoryId
}

func (m Movie) Director() DirectorId {
	return m.director
}

func (m Movie) DisplayName() DisplayName {
	return m.displayName
}

func (m Movie) Description() Description {
	return m.description
}

func (m Movie) Year() Year {
	return m.year
}

func (m Movie) Picture() Picture {
	return m.picture
}

func (m Movie) WatchUrl() WatchUrl {
	return m.watchUrl
}

func (m Movie) CrawlUrl() CrawlUrl {
	return m.crawlUrl
}
