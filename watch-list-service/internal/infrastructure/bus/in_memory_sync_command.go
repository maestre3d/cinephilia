package bus

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/neutrinocorp/ddderr"
)

//	@Adapter
//	@Bus
//	@Sync
type InMemorySyncCommand struct {
	handlerMap map[string]domain.CommandHandler
}

func NewInMemorySyncCommand() *InMemorySyncCommand {
	return &InMemorySyncCommand{handlerMap: map[string]domain.CommandHandler{}}
}

func (c *InMemorySyncCommand) RegisterHandler(command domain.Command, handler domain.CommandHandler) error {
	if _, ok := c.handlerMap[command.Id()]; ok {
		return ddderr.NewAlreadyExists(nil, "command")
	}

	c.handlerMap[command.Id()] = handler
	return nil
}

func (c InMemorySyncCommand) Dispatch(ctx context.Context, command domain.Command) error {
	if _, ok := c.handlerMap[command.Id()]; !ok {
		return ddderr.NewNotFound(nil, "command")
	}

	return c.handlerMap[command.Id()].Invoke(ctx, command)
}
