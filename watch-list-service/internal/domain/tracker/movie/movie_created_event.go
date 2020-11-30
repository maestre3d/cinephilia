package movie

// MovieCreated A movie was created successfully
//	@DomainEvent
type MovieCreated struct {
	id          string
	name        string
	description string
}

func newMovieCreated(id, name, description string) MovieCreated {
	return MovieCreated{
		id:          id,
		name:        name,
		description: description,
	}
}

func (_ MovieCreated) Name() string {
	return "movie.created"
}

func (c MovieCreated) ToPrimitive() map[string]string {
	return map[string]string{
		"movie_id":     c.id,
		"display_name": c.name,
		"description":  c.description,
	}
}
