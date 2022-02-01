package routes

import (
	"freecharge/rsrc-bp/api/resources/health"

	"github.com/gofiber/fiber/v2"
)

func (r *RouteHandler) Health(a fiber.Router) {
	handler := health.NewHealthHandler(r.sConfig.ProductName, r.sConfig.ModuleName)

	a.Get("/", handler.Health)
}
