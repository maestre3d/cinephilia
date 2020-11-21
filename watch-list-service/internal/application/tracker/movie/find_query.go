package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// FindQuery requests a movie
//	@DTO
//	@Query
type FindQuery struct {
	MovieId string
}

func (q FindQuery) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.movie.find"
}
