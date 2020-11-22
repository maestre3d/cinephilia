package marshaler

import (
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
)

// MoviePrimitive movie anemic model
//	@Adapter
//	@DTO
type MoviePrimitive struct {
	Id          string
	UserId      string
	DisplayName string
	Description *string
	Year        *uint32
	Picture     *string
	WatchUrl    *string
	CrawlUrl    *string
	CreateTime  time.Time
	UpdateTime  time.Time
	Active      bool
}

func (s MoviePrimitive) UnmarshalAggregate() (*movie.Movie, error) {
	var resultErr *multierror.Error
	id, err := movie.NewMovieId(s.Id)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	userId, err := movie.NewUserId(s.UserId)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	displayName, err := movie.NewDisplayName(s.DisplayName)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	description, err := movie.NewDescription(s.SafeDescription())
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	year, err := movie.NewYear(s.SafeYear())
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	picture, err := movie.NewPicture(s.SafePicture())
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	watchUrl, err := movie.NewWatchUrl(s.SafeWatchUrl())
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	crawlUrl, err := movie.NewCrawlUrl(s.SafeCrawlUrl())
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
	}
	if resultErr != nil && resultErr.ErrorOrNil() != nil {
		return nil, resultErr
	}

	return movie.NewMovie(movie.CreateArgs{
		Id:          *id,
		UserId:      *userId,
		DisplayName: *displayName,
		Description: *description,
		Year:        *year,
		Picture:     *picture,
		WatchUrl:    *watchUrl,
		CrawlUrl:    *crawlUrl,
	}), nil
}

func (s MoviePrimitive) SafeDescription() string {
	if s.Description == nil {
		return ""
	}

	return *s.Description
}

func (s MoviePrimitive) SafeYear() uint32 {
	if s.Year == nil {
		return 0
	}

	return *s.Year
}

func (s MoviePrimitive) SafePicture() string {
	if s.Picture == nil {
		return ""
	}

	return *s.Picture
}

func (s MoviePrimitive) SafeWatchUrl() string {
	if s.WatchUrl == nil {
		return ""
	}

	return *s.WatchUrl
}

func (s MoviePrimitive) SafeCrawlUrl() string {
	if s.CrawlUrl == nil {
		return ""
	}

	return *s.CrawlUrl
}
