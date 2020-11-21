package domain

import "context"

// CommandBus receives Command(s) and dispatch them to their respective command handler
//	@Port
//	@Bus
//	@Command
type CommandBus interface {
	RegisterHandler(command Command, handler CommandHandler) error
	Dispatch(ctx context.Context, command Command) error
}
