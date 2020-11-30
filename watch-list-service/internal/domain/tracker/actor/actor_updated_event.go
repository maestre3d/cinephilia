package actor

// ActorUpdated An actor was created successfully
//	@DomainEvent
type ActorUpdated struct {
	id      string
	userId  string
	name    string
	picture string
}

func newActorUpdated(id, userId, name, picture string) ActorUpdated {
	return ActorUpdated{
		id:      id,
		userId:  userId,
		name:    name,
		picture: picture,
	}
}

func (_ ActorUpdated) Name() string {
	return "actor.updated"
}

func (u ActorUpdated) ToPrimitive() map[string]string {
	return map[string]string{
		"actor_id":     u.id,
		"user_id":      u.userId,
		"display_name": u.name,
		"picture":      u.picture,
	}
}
