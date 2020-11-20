package movie

import "github.com/neutrinocorp/ddderr"

var (
	UserRequired = ddderr.NewRequired("movie_user_id")

	NameRequired       = ddderr.NewRequired("movie_display_name")
	NameAboveMaxLength = ddderr.NewOutOfRange("movie_display_name", "1", "128")

	DescriptionAboveMaxLength = ddderr.NewOutOfRange("movie_description", "0", "512")

	InvalidPictureExtension = ddderr.NewInvalidFormat("movie_picture", "jpg, jpeg, webp and png")

	AlreadyExists = ddderr.NewAlreadyExists(nil, "movie")
	NotExists     = ddderr.NewNotFound(nil, "movie")
)
