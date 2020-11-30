package domain

// Event domain event, represents an state change of an aggregate within the ecosystem
//	@DTO
type Event interface {
	Name() string
	ToPrimitive() map[string]string
}
