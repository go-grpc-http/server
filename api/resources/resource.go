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
			Status:     http.StatusText(http.StatusInternalServerError),
		}

		// TODO: need to set the trace tree
		eResp.Error.Message = err.Error()

		c.Status(http.StatusInternalServerError).JSON(eResp)
	}
	return nil
}

func ErrorWrapper(fn Handler) fiber.Handler {
	return Handler(fn).Error
}
