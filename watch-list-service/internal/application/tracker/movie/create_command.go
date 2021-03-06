package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// CreateCommand requests a Movie creation
//	@DTO
//	@Command
type CreateCommand struct {
	MovieId     string
	UserId      string
	DisplayName string
	Description string
	Year        string
	Picture     string
	WatchUrl    string
	CrawlUrl    string
}

func (c CreateCommand) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.movie.create"
}
