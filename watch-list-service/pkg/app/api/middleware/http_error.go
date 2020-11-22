package middleware

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/neutrinocorp/ddderr"
)

func ErrorHTTP(c *fiber.Ctx) error {
	err := c.Next()
	switch {
	case errors.Is(err, ddderr.AlreadyExists):
		return fiber.NewError(http.StatusConflict, ddderr.GetDescription(err))
	case errors.Is(err, ddderr.NotFound):
		return fiber.NewError(http.StatusNotFound, ddderr.GetDescription(err))
	case errors.Is(err, ddderr.FailedRemoteCall):
		return fiber.NewError(http.StatusServiceUnavailable, ddderr.GetDescription(err))
	case ddderr.IsDomain(err):
		return fiber.NewError(http.StatusBadRequest, ddderr.GetDescription(err))
	case ddderr.IsInfrastructure(err):
		return fiber.NewError(http.StatusInternalServerError, ddderr.GetDescription(err))
	default:
		return err
	}
}

func ErrorHandlerHTTP(ctx *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Retreive the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).JSON(&fiber.Map{
		"error":  err.Error(),
		"status": code,
	})
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(500).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
