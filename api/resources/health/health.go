package health

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	projectName string
	modelName   string
}

func NewHealthHandler(pName, mName string) *HealthHandler {
	return &HealthHandler{
		projectName: pName,
		modelName:   mName,
	}
}

// Health - generic health check
func (h *HealthHandler) Health(c *fiber.Ctx) error {
	response := HealthResponse{
		StatusCode: 200,
		Status:     "OK",
		Message:    fmt.Sprintf("health check is successful for: %s", h.projectName),
	}
	err := c.Status(http.StatusOK).JSON(response)
	if err != nil {
		return err
	}
	return errors.New("testing error")
	// return nil
}
