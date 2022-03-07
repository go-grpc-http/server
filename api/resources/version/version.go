package version

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/middleware/libs/response"
)

type VersionHandler struct {
	projectName string
	modelName   string
	version     string
}

func NewVersionHandler(pName, mName, version string) *VersionHandler {
	return &VersionHandler{
		projectName: pName,
		modelName:   mName,
		version:     version,
	}
}

// generic version check
func (h *VersionHandler) Version(c *fiber.Ctx) error {
	rData := VersionResponseData{
		Version:     h.version,
		ModelName:   h.modelName,
		ProjectName: h.projectName,
	}

	return response.NewBody(c, http.StatusOK, "version check successful", rData, nil)
}
