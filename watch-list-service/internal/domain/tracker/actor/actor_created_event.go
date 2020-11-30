package actor

// ActorCreated An actor was created successfully
//	@DomainEvent
type ActorCreated struct {
	id      string
	userId  string
	name    string
	picture string
}

func newActorCreated(id, userId, name, picture string) ActorCreated {
	return ActorCreated{
		id:      id,
		userId:  userId,
		name:    name,
		picture: picture,
	}
}

func (_ ActorCreated) Name() string {
	return "actor.created"
}

func (c ActorCreated) ToPrimitive() map[string]string {
	return map[string]string{
		"actor_id":     c.id,
		"user_id":      c.userId,
		"display_name": c.name,
		"picture":      c.picture,
	}
}
