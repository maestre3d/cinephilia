package querybus

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"go.uber.org/zap"
)

func NewQueryBus(bus domain.QueryBus, logger *zap.Logger) domain.QueryBus {
	return log{
		logger: logger,
		next:   bus,
	}
}
