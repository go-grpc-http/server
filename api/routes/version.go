package routes

import (
	"freecharge/rsrc-bp/api/resources/version"

	"github.com/gofiber/fiber/v2"
)

func (r *RouteHandler) Version(a fiber.Router) {
	handler := version.NewVersionHandler(r.sConfig.ProductName, r.sConfig.ModuleName, r.sConfig.Version)

	a.Get("/", handler.Version)
}
