package version

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/middleware/libs/response"
)

// Version just a version check
func (h *Config) Version(c *fiber.Ctx) error {
	rData := VersionResponseData{
		Version:     h.version,
		ModelName:   h.modelName,
		ProjectName: h.projectName,
	}

	return response.NewBody(c, http.StatusOK, "version check successful", rData, nil)
}
