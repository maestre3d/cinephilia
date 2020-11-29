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

func (c DirectorCreated) GetName() string {
	return "director.created"
}

func (c DirectorCreated) ToPrimitive() interface{} {
	return nil
}
