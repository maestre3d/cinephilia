package logging

import (
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/configuration"
	"go.uber.org/zap"
)

//	@Adapter
type Zap struct {
	log *zap.Logger
}

func NewZapPrimitive(cfg configuration.Configuration) (logger *zap.Logger, err error) {
	if cfg.IsProdEnv() || cfg.IsStagingEnv() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

func NewZap() (domain.Logger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	logger := &Zap{log: zapLogger}
	return logger, nil
}

func (z Zap) parseFields(args []domain.LogField) []zap.Field {
	fields := make([]zap.Field, 0)
	for _, a := range args {
		fields = append(fields, zap.Any(a.Key, a.Value))
	}
	return fields
}

func (z *Zap) Debug(msg string, args ...domain.LogField) {
	z.log.Debug(msg, z.parseFields(args)...)
}

func (z *Zap) Info(msg string, args ...domain.LogField) {
	z.log.Info(msg, z.parseFields(args)...)
}

func (z *Zap) Warn(msg string, args ...domain.LogField) {
	z.log.Warn(msg, z.parseFields(args)...)
}

func (z *Zap) Error(msg string, args ...domain.LogField) {
	z.log.Error(msg, z.parseFields(args)...)
}
