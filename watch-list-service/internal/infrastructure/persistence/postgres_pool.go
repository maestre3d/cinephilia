package persistence

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	pqConn      *sql.DB
	pqSingleton *sync.Once
)

func init() {
	pqSingleton = new(sync.Once)
}

func NewPostgresPool() (*sql.DB, error) {
	var err error
	pqSingleton.Do(func() {
		connString := "postgres://postgres:postgres@localhost:5432/watch_list?sslmode=disable"
		pqConn, err = sql.Open("postgres", connString)
	})

	return pqConn, err
}
