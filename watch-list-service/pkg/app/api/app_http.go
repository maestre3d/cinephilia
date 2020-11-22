package api

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	movieapp "github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/bus"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/persistence"
	movieinfra "github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie"
	movpersistence "github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/persistence"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/webscrap"
	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api/controller"
	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api/middleware"
	"go.uber.org/fx"
)

var httpApp *fiber.App

func InitHTTP(ctx context.Context) *fx.App {
	return fx.New(
		fx.Provide(
			newRouter,
			persistence.NewPostgresPool,
			func(pool *sql.DB) movie.Repository {
				return movpersistence.NewPostgresMovieRepository(pool)
			},
			func(repository movie.Repository) domain.CommandBus {
				return bus.NewInMemorySyncCommand()
			},
			func() domain.QueryBus {
				return bus.NewInMemorySyncQuery()
			},
			movieapp.NewCreator,
			movieapp.NewCreateCommandHandler,
			webscrap.NewCollyImdbCollector,
			movieinfra.NewImdbWebScrapper,
			func(scrapper *movieinfra.ImdbWebScrapper) movie.MovieCrawler {
				return movieinfra.NewImdbMovieCrawler(scrapper)
			},
			movieapp.NewCreateByCrawlCommandHandler,
			movieapp.NewFinder,
			movieapp.NewFindQueryHandler,
		),
		fx.Invoke(
			func(commandBus domain.CommandBus, createCrawl *movieapp.CreateByCrawlCommandHandler,
				create *movieapp.CreateCommandHandler) error {
				err := commandBus.RegisterHandler(movieapp.CreateByCrawlCommand{},
					createCrawl)
				if err != nil {
					return err
				}

				return commandBus.RegisterHandler(movieapp.CreateCommand{}, create)
			},
			func(queryBus domain.QueryBus, find *movieapp.FindQueryHandler) (domain.QueryBus, error) {
				err := queryBus.RegisterHandler(movieapp.FindQuery{}, find)
				return queryBus, err
			},
			controller.NewMovieHTTP,
			func() error {
				return httpApp.Listen(":8081")
			},
		),
	)
}

func newRouter(lc fx.Lifecycle) fiber.Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandlerHTTP,
	})
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(compress.New())
	//app.Use(csrf.New())
	app.Use(middleware.XssHTTP)
	app.Use(recover.New())
	app.Use(middleware.ErrorHTTP)
	app.Get("/", func(c *fiber.Ctx) error {
		_ = c.SendString("Welcome to Watch List API")
		return nil
	})
	api := app.Group("/v1")
	httpApp = app

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpApp.Shutdown()
		},
	})

	return api
}
