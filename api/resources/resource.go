package resources

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler func(c *fiber.Ctx) error

func (fn Handler) Error(c *fiber.Ctx) {
	if err := fn(c); err != nil {
		eResp := &ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Error:      err,
		}
		c.Status(http.StatusInternalServerError).JSON(eResp)
	}
}
