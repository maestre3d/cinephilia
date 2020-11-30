package controller

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
)

// MovieHTTP handles Movie http endpoints
//	@Controller
//	@API
//	@HTTP
type MovieHTTP struct {
	commandBus domain.CommandBus
	queryBus   domain.QueryBus
}

func NewMovieHTTP(router fiber.Router, commandBus domain.CommandBus, queryBus domain.QueryBus) *MovieHTTP {
	mov := &MovieHTTP{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
	mov.initRouting(router)
	return mov
}

func (h MovieHTTP) initRouting(router fiber.Router) {
	router.Get("/movies/:id", h.getMovie)
	router.Put("/users/:user/movies/:movie", h.createMovie)
	router.Put("/users/:user/movies/:movie/crawl", h.requestCrawlMovie)
	router.Get("/users/:user/movies", h.listUserMovies)
}

//	@Get("id")
func (h MovieHTTP) getMovie(c *fiber.Ctx) error {
	mov, err := h.queryBus.Ask(c.Context(), movie.FindQuery{MovieId: c.Params("id")})
	if err != nil {
		return err
	}

	return c.JSON(mov)
}

//	@Get("user_id")
func (h MovieHTTP) listUserMovies(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("implement %s's movie listing", c.Params("user")))
}

//	@Put("user_id", "movie_id")
func (h MovieHTTP) createMovie(c *fiber.Ctx) error {
	err := h.commandBus.Dispatch(c.Context(), movie.CreateCommand{
		MovieId:     c.Params("movie"),
		DisplayName: c.FormValue("display_name"),
		UserId:      c.Params("user"),
		Description: c.FormValue("description"),
		Year:        c.FormValue("year"),
		WatchUrl:    c.FormValue("watch_url"),
		Picture:     c.FormValue("picture"),
	})
	if err != nil {
		return err
	}

	c.Status(http.StatusCreated)
	return nil
}

//	@Put("user_id", "movie_id")
func (h MovieHTTP) requestCrawlMovie(c *fiber.Ctx) error {
	err := h.commandBus.Dispatch(c.Context(), movie.CreateByCrawlCommand{
		MovieId:  c.Params("movie"),
		UserId:   c.Params("user"),
		CrawlUrl: c.FormValue("crawl_url"),
	})
	if err != nil {
		return err
	}

	c.Status(http.StatusAccepted)
	return nil
}
