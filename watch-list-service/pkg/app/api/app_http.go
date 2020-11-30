package api

import (
	"context"
	"database/sql"
	"time"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/bus/commandbus"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/bus/querybus"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/configuration"

	"go.uber.org/zap"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/logging"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	movieapp "github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/persistence"
	movieinfra "github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie"
	movpersistence "github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/persistence"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/webscrap"
	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api/controller"
	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api/middleware"
	"go.uber.org/fx"
)

func InitHTTP(_ context.Context) *fx.App {
	return fx.New(
		fx.Provide(
			configuration.NewConfiguration,
			logging.NewZapPrimitive,
			persistence.NewPostgresPool,
			func(pool *sql.DB) movie.Repository {
				return movpersistence.NewPostgresMovieRepository(pool)
			},
			func(log *zap.Logger) domain.CommandBus {
				return commandbus.NewCommandBus(commandbus.NewInMemorySyncCommand(), log)
			},
			func(log *zap.Logger) domain.QueryBus {
				return querybus.NewQueryBus(querybus.NewInMemorySyncQuery(), log)
			},
			movieapp.NewCreator,
			webscrap.NewCollyImdbCollector,
			movieinfra.NewImdbWebScrapper,
			func(scrapper *movieinfra.ImdbWebScrapper) movie.MovieCrawler {
				return movieinfra.NewImdbMovieCrawler(scrapper)
			},
			movieapp.NewFinder,
			newHTTPApp,
			newHTTPRouter,
		),
		fx.Invoke(
			movieapp.NewCreateCommandHandler,
			movieapp.NewCreateByCrawlCommandHandler,
			movieapp.NewFindQueryHandler,
			controller.NewHealthCheckHTTP,
			controller.NewMovieHTTP,
			func(log *zap.Logger, cfg configuration.Configuration) {
				log.Info("starting http application", zap.Namespace("config"),
					zap.String("service", cfg.Service),
					zap.String("stage", cfg.Stage),
					zap.String("version", cfg.Version),
				)
			},
			startHTTP,
		),
	)
}

func newHTTPApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandlerHTTP,
		ReadTimeout:  time.Second * 5,
	})
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(compress.New())
	//app.Use(csrf.New())
	app.Use(middleware.XssHTTP)
	app.Use(recover.New())
	app.Use(middleware.ErrorHTTP)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Watch List API")
	})

	return app
}

func newHTTPRouter(app *fiber.App) fiber.Router {
	return app.Group("/v1")
}

func startHTTP(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) (err error) {
			go func() {
				err = app.Listen(":8080")
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
