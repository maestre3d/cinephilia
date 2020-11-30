package director

// DirectorCreated A director was created successfully
//	@DomainEvent
type DirectorCreated struct {
	id      string
	userId  string
	name    string
	picture string
}

func NewDirectorCreated(id, userId, name, picture string) DirectorCreated {
	return DirectorCreated{
		id:      id,
		userId:  userId,
		name:    name,
		picture: picture,
	}
}

func (c DirectorCreated) Name() string {
	return "director.created"
}

func (c DirectorCreated) ToPrimitive() map[string]string {
	return map[string]string{
		"director_id":  c.id,
		"user_id":      c.userId,
		"display_name": c.name,
		"picture":      c.picture,
	}
}
