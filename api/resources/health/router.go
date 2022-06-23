package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/logger"
	"github.com/rohanraj7316/rsrc-bp-testing/configs"
)

func Router(a fiber.Router) {
	// load the configs
	cfg, err := configs.NewServerConfig()
	if err != nil {
		logger.Error(err.Error())
	}

	// initialize your handler
	handler := New(cfg.ProductName, cfg.ModuleName)

	// declare your routes
	a.Get("/", handler.Health)
}
