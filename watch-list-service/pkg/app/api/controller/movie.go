package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/neutrinocorp/ddderr"
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
	router.Get("/movie/:id", h.getMovie)
}

//	@Get("id")
func (h MovieHTTP) getMovie(c *fiber.Ctx) error {
	mov, err := h.queryBus.Ask(c.Context(), movie.FindQuery{MovieId: c.Params("id")})
	if err != nil {
		_ = c.JSON(&fiber.Map{
			"success": false,
			"error":   ddderr.GetDescription(err),
		})
		return err
	}

	_ = c.JSON(mov)
	return nil
}
