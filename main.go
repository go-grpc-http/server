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
	// can be used to terminate the server using done
	pCtx := context.Background()
	ctx, cancel := context.WithCancel(pCtx)
	defer cancel()

	// initialize logger
	err := logger.Configure()
	if err != nil {
		log.Panic(err)
	}

	config, err := configs.NewServerConfig()
	if err != nil {
		logger.Error(err.Error())
		cancel()
	}

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
		logger.Error(err.Error())
		cancel()
	}

	r.NewRouter(app)

	logger.Info("successful starting server :)")

	err = app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Error(err.Error())
		cancel()
	}

	cChannel := make(chan os.Signal, 2)
	signal.Notify(cChannel, os.Interrupt, syscall.SIGTERM)

bLoop:
	for {
		select {
		case <-ctx.Done():
			break bLoop

		case <-cChannel:
			logger.Warn("catch interrupted signal")
			time.Sleep(config.WaitTimeBeforeKill)
			break bLoop
		}
	}

	// TODO: add logics for service shutdown.
	err = app.Shutdown()
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Warn("shutting down the server :(")
}
