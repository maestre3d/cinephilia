package movie

import "github.com/neutrinocorp/ddderr"

var (
	NameRequired       = ddderr.NewRequired("movie_display_name")
	NameAboveMaxLength = ddderr.NewOutOfRange("movie_display_name", "1", "128")

	DescriptionAboveMaxLength = ddderr.NewOutOfRange("movie_description", "0", "512")

	MovieAlreadyExists = ddderr.NewAlreadyExists(nil, "movie")
)
