package domain

import "context"

// EventPublisher pushes a domain Event into the event bus
//	@Port
type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}
