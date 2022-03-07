package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rohanraj7316/rsrc-bp-testing/api/routes"
	"github.com/rohanraj7316/rsrc-bp-testing/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"github.com/rohanraj7316/logger"
	"github.com/rohanraj7316/middleware"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// TODO: need to check for the cancelation of contexts
	// can be used to terminate the server using done
	pCtx := context.Background()
	ctx, cancel := context.WithCancel(pCtx)
	defer cancel()

	// initialize logger
	lOptions, _ := configs.NewLogConfig(logger.NewOptions())
	err := logger.Configure(lOptions)
	if err != nil {
		log.Panic(err)
	}

	config := configs.NewServerConfig(cancel)

	app := fiber.New(config.ServerConfig)

	// adding middleware
	app.Use(cors.New(config.CorsConfig))
	app.Use(helmet.New())

	// initializing middleware
	/*
		used to turn off the request logging
		middleware.ConfigDefault.SetReqResLog(false)
		used to turn off request & response body logging
		middleware.ConfigDefault.SetReqResBodyLog(false)
	*/
	app.Use(middleware.New(app))

	// initialize router
	r, err := routes.NewRouteHandler(app, config)
	if err != nil {
		cancel()
	}

	r.NewRouter(app)

	logger.Info("successful starting server :)")
	// TODO: add request timeout and other configs to the server
	err = app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		cancel()
	}

	cChannel := make(chan os.Signal, 2)
	signal.Notify(cChannel, os.Interrupt, syscall.SIGTERM)

bLoop:
	for {
		select {
		case <-ctx.Done():
			// TODO: need to understand this more deaply.
			// list down the operations after Done is triggered

		case <-cChannel:
			logger.Warn("catch interrupted signal")
			time.Sleep(config.WaitTimeBeforeKill)
			break bLoop
		}
	}

	// TODO: add grpc shutting down handling
	err = app.Shutdown()
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Warn("shutting down the server :(")
}
