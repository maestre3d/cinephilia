package movie

// CreateCommand requests a Movie creation
//	@DTO
//	@Command
type CreateCommand struct {
	Id          string
	DisplayName string
	Description string
	UserId      string
}
