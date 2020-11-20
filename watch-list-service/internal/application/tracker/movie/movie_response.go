package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"

// MovieResponse query response for one movie
//	@DTO
type MovieResponse struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	CategoryId  *string `json:"category_id"`
	DisplayName string  `json:"display_name"`
	Description *string `json:"description"`
	Picture     *string `json:"picture"`
	WatchUrl    *string `json:"watch_url"`
	CrawlUrl    *string `json:"crawl_url"`
}

func newMovieResponseFromAggregate(mov movie.Movie) *MovieResponse {
	return &MovieResponse{
		Id:          mov.Id().Value(),
		UserId:      mov.User().Value(),
		CategoryId:  mov.Category().NullableValue(),
		DisplayName: mov.DisplayName().Value(),
		Description: mov.Description().NullableValue(),
		Picture:     mov.Picture().NullableValue(),
		WatchUrl:    mov.WatchUrl().NullableValue(),
		CrawlUrl:    mov.CrawlUrl().NullableValue(),
	}
}
