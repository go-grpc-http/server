package routes

import (
	"fmt"
	"net/http"

	"github.com/rohanraj7316/middleware/libs/response"
	"github.com/rohanraj7316/rsrc-bp-testing/configs"

	"github.com/gofiber/fiber/v2"
)

type Router func(fiber.Router)

type Route struct {
	path   string
	router Router
}

type RouteHandler struct {
	app     *fiber.App
	sConfig *configs.ServerConfigStruct
}

func NewRouteHandler(app *fiber.App, sConfig *configs.ServerConfigStruct) (*RouteHandler, error) {
	return &RouteHandler{
		app:     app,
		sConfig: sConfig,
	}, nil
}

func (r *RouteHandler) NewRouter(app *fiber.App) {
	// list down all the routes and their handlers
	routes := []Route{
		{
			path:   "/health",
			router: r.Health,
		},
		{
			path:   "/version",
			router: r.Version,
		},
	}

	for i := 0; i < len(routes); i++ {
		route := routes[i]
		aGroup := app.Group(route.path)
		route.router(aGroup)
	}

	app.Use("*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Cannot %s %s", c.Method(), c.Path()) // Cannot GET /healths
		return response.NewBody(c, http.StatusInternalServerError, msg, nil, nil)
	})
}
