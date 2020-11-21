package domain

// Command request an action
//	@DTO
//	@Command
type Command interface {
	Id() string
}
