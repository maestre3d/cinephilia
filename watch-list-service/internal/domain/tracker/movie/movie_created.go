package movie

// MovieCreated A movie was created successfully
//	@DomainEvent
type MovieCreated struct {
	id          string
	name        string
	description string
}

func NewMovieCreated(id, name, description string) MovieCreated {
	return MovieCreated{
		id:          id,
		name:        name,
		description: description,
	}
}

func (c MovieCreated) GetName() string {
	return "movie.created"
}

func (c MovieCreated) ToPrimitive() interface{} {
	return nil
}
