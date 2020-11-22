package persistence

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"go.uber.org/fx"

	_ "github.com/lib/pq"
)

var (
	pqConn      *sql.DB
	pqSingleton *sync.Once
)

func init() {
	pqSingleton = new(sync.Once)
}

func NewPostgresPool(lc fx.Lifecycle) (*sql.DB, error) {
	var err error
	pqSingleton.Do(func() {
		connString := "postgres://postgres:postgres@localhost:5432/watch_list?sslmode=disable"
		pqConn, err = sql.Open("postgres", connString)
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if pqConn != nil {
				log.Print("closing postgres")
				return pqConn.Close()
			}
			return nil
		},
	})

	return pqConn, err
}
