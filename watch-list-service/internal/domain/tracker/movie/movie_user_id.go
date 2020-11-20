package movie

import "github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker"

// UserId user unique identifier, assigns a movie to a user
//	@ValueObject
type UserId struct {
	// extends userId
	value tracker.UserId
}

func NewUserId(value string) (*UserId, error) {
	userId := new(UserId)
	if err := userId.ensureRequired(value); err != nil {
		return nil, err
	}
	id, err := tracker.NewUserId("movie_user_id", value)
	if err != nil {
		return nil, err
	}

	return &UserId{value: *id}, nil
}

func (i UserId) ensureRequired(value string) error {
	//	rules
	//	a.	value.length > 0
	if value == "" {
		return UserRequired
	}

	return nil
}

func (i UserId) Value() string {
	return i.value.Value()
}
