package health

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/middleware/libs/response"
)

type HealthHandler struct {
	projectName string
	modelName   string
}

func New(pName, mName string) *HealthHandler {
	return &HealthHandler{
		projectName: pName,
		modelName:   mName,
	}
}

// Health just a health check
func (h *HealthHandler) Health(c *fiber.Ctx) error {
	return response.NewBody(c, http.StatusOK, fmt.Sprintf("health check is successful for: %s", h.projectName), nil, nil)
}
