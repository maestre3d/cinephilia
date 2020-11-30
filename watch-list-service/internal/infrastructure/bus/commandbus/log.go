package commandbus

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/neutrinocorp/ddderr"
	"go.uber.org/zap"
)

type log struct {
	logger *zap.Logger
	next   domain.CommandBus
}

func (l log) RegisterHandler(command domain.Command, handler domain.CommandHandler) (err error) {
	defer func() {
		if err != nil {
			l.logger.Error("failed to register command",
				zap.String("command", command.Id()),
			)
			return
		}

		l.logger.Info("registered command", zap.String("command", command.Id()))
	}()

	err = l.next.RegisterHandler(command, handler)
	return err
}

func (l log) Dispatch(ctx context.Context, command domain.Command) (err error) {
	defer func() {
		if err != nil {
			l.logger.Warn("failed to dispatch command",
				zap.String("command", command.Id()),
				zap.String("desc", ddderr.GetDescription(err)),
				zap.String("parent_desc", ddderr.GetParentDescription(err)),
				zap.Error(err),
			)
			return
		}

		l.logger.Info("dispatched command", zap.String("command", command.Id()))
	}()

	err = l.next.Dispatch(ctx, command)
	return err
}
