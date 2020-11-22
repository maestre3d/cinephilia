package movie

import "github.com/neutrinocorp/ddderr"

var (
	IdRequired   = ddderr.NewRequired("movie_id")
	UserRequired = ddderr.NewRequired("movie_user_id")

	NameRequired = ddderr.NewRequired("movie_display_name")

	DescriptionAboveMaxLength = ddderr.NewOutOfRange("movie_description", "0", "512")

	YearAboveMaxLength = ddderr.NewOutOfRange("movie_year", "0", "3000")
	YearBelowMinLength = ddderr.NewOutOfRange("movie_year", "0", "3000")

	CrawlUrlIsNotAvailable = ddderr.NewInvalidFormat("movie_crawl_url", "IMDb link")

	AlreadyExists = ddderr.NewAlreadyExists(nil, "movie")
	NotExists     = ddderr.NewNotFound(nil, "movie")
)
