package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/configuration"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var (
	pqConn      *sql.DB
	pqSingleton *sync.Once
)

func init() {
	pqSingleton = new(sync.Once)
}

func NewPostgresPool(lc fx.Lifecycle, logger *zap.Logger, cfg configuration.Configuration) (*sql.DB, error) {
	var err error
	pqSingleton.Do(func() {
		logger.With(
			zap.Namespace("metadata"),
			zap.String("database", cfg.Postgres.Database),
			zap.String("address", cfg.Postgres.Address),
			zap.Int("port", cfg.Postgres.Port),
			zap.Bool("secure_mode", cfg.Postgres.SecureMode),
		).Info("connecting to postgresql")

		// "postgres://postgres:postgres@localhost:5432/watch_list?sslmode=disable"
		pqConn, err = sql.Open("postgres", generatePostgresConnString(cfg))
		if err != nil {
			logger.Error("could not connect to postgresql", zap.Error(err))
		}
	})

	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			if pqConn != nil {
				logger.Info("disconnecting from postgresql")
				return pqConn.Close()
			}
			return nil
		},
	})

	return pqConn, err
}

func generatePostgresConnString(cfg configuration.Configuration) string {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Address,
		cfg.Postgres.Port, cfg.Postgres.Database)
	if !cfg.Postgres.SecureMode {
		connString += "?sslmode=disable"
	}

	return connString
}
