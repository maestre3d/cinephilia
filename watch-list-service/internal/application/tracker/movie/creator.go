package movie

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

// Creator creates a movie using basic data
//	@ApplicationService
type Creator struct {
	repository movie.Repository
	eventBus   domain.EventBus
}

//	@DTO
type CreatorArgs struct {
	Id          movie.Id
	UserId      movie.UserId
	CategoryId  movie.CategoryId
	DirectorId  movie.DirectorId
	DisplayName movie.DisplayName
	Description movie.Description
	Year        movie.Year
	Picture     movie.Picture
	WatchUrl    movie.WatchUrl
	CrawlUrl    movie.CrawlUrl
}

func NewCreator(repository movie.Repository) *Creator {
	return &Creator{repository: repository}
}

func (c Creator) Create(ctx context.Context, args CreatorArgs) error {
	if mov, _ := NewFinder(c.repository).Find(ctx, args.Id); mov != nil {
		return movie.AlreadyExists
	}

	return c.repository.Save(ctx, *movie.NewMovie(movie.CreateArgs{
		Id:          args.Id,
		UserId:      args.UserId,
		CategoryId:  args.CategoryId,
		DirectorId:  args.DirectorId,
		DisplayName: args.DisplayName,
		Description: args.Description,
		Year:        args.Year,
		Picture:     args.Picture,
		WatchUrl:    args.WatchUrl,
		CrawlUrl:    args.CrawlUrl,
	}))
}
