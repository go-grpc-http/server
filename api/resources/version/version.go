package version

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	response := VersionResponse{
		StatusCode: 200,
		Status:     "OK",
		Message:    "version check successful",
		Data: Data{
			Version:     h.version,
			ModelName:   h.modelName,
			ProjectName: h.projectName,
		},
	}
	err := c.Status(http.StatusOK).JSON(response)
	if err != nil {
		return err
	}
	return nil
}
