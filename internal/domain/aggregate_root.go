package domain

// AggregateRoot a domain module, set of value objects and an entity which handles part of the
// business rules and operations
//	@Aggregate
type AggregateRoot interface {
	record(e ...Event)
	PullEvents() []Event
}
