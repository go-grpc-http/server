package resources

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler func(c *fiber.Ctx) error

func (fn Handler) Error(c *fiber.Ctx) error {
	if err := fn(c); err != nil {
		eResp := &ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Error:      err,
		}
		c.Status(http.StatusInternalServerError).JSON(eResp)
	}
	return nil
}

func ErrorWrapper(fn Handler) fiber.Handler {
	return Handler(fn).Error
}
