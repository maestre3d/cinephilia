package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

//	@DTO
type CreateArgs struct {
	Id          Id
	UserId      UserId
	DisplayName DisplayName
	Description Description
	Year        Year
	Picture     Picture
	WatchUrl    WatchUrl
	CrawlUrl    CrawlUrl
}

func NewMovie(args CreateArgs) *Movie {
	mov := &Movie{
		id:          args.Id,
		user:        args.UserId,
		displayName: args.DisplayName,
		description: args.Description,
		year:        args.Year,
		picture:     args.Picture,
		watchUrl:    args.WatchUrl,
		crawlUrl:    args.CrawlUrl,
		events:      make([]domain.Event, 0),
	}
	mov.record(newMovieCreated(mov.Id().Value(), mov.DisplayName().Value(), mov.Description().Value()))
	return mov
}
