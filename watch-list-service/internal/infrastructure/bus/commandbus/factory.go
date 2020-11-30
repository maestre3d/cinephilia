package commandbus

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"go.uber.org/zap"
)

func NewCommandBus(bus domain.CommandBus, logger *zap.Logger) domain.CommandBus {
	return log{
		logger: logger,
		next:   bus,
	}
}
