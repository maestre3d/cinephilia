package persistence

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/marshaler"
)

//	@Repository
//	@Adapter
type PostgresMovieRepository struct {
	db *sql.DB
	mu *sync.RWMutex
}

func NewPostgresMovieRepository(db *sql.DB) *PostgresMovieRepository {
	return &PostgresMovieRepository{db: db, mu: new(sync.RWMutex)}
}

func (p *PostgresMovieRepository) Save(ctx context.Context, movie movie.Movie) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	conn, err := p.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `INSERT INTO movie(id, user_id, display_name, description, year, picture, watch_url, 
                  crawl_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err = conn.ExecContext(ctx, query, movie.Id().Value(), movie.User().Value(),
		movie.DisplayName().Value(), movie.Description().NullableValue(), movie.Year().NullableValue(),
		movie.Picture().NullableValue(), movie.WatchUrl().NullableValue(), movie.CrawlUrl().NullableValue())
	return err
}

func (p *PostgresMovieRepository) Search(ctx context.Context, movieId movie.Id) (*movie.Movie, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	conn, err := p.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	movSQL := new(marshaler.MoviePrimitive)
	err = conn.QueryRowContext(ctx, `SELECT * FROM movie WHERE id = $1 AND active = TRUE;`, movieId.Value()).
		Scan(&movSQL.Id, &movSQL.UserId, &movSQL.DisplayName, &movSQL.Description, &movSQL.Year,
			&movSQL.Picture, &movSQL.WatchUrl, &movSQL.CreateTime, &movSQL.UpdateTime, &movSQL.Active,
			&movSQL.CrawlUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return movSQL.UnmarshalAggregate()
}

func (p *PostgresMovieRepository) SearchByCriteria(ctx context.Context, criteria domain.Criteria) ([]*movie.Movie, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return nil, nil
}
