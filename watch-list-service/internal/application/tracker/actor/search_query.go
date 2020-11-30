package actor

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

// SearchQuery request an actors list
//	@DTO
//	@Query
type SearchQuery struct {
	PageSize  int32
	PageToken string
	OrderBy   string
	OrderKind string
	Filters   []map[string]string
}

func (q SearchQuery) Id() string {
	return domain.CinephiliaDomainAlt + ".tracker.actor.search"
}
