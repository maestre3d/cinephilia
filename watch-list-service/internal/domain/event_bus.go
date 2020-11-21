package domain

import "context"

// EventBus pushes a domain Event into the event bus, subscribes to domain events
//	@Port
type EventBus interface {
	Publish(ctx context.Context, events ...Event) error
}
