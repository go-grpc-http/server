package main

import (
	"context"
	"fmt"
	"freecharge/rsrc-bp/api/helpers/logger"
	"freecharge/rsrc-bp/configs"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	config := configs.NewServerConfig(cancel)
	sConfig := fiber.Settings{}

	cConfig :=  cors.Config{
		AllowOrigins: sConfig.CorsConfig.AllowOrigins
		AllowMethods: sConfig.CorsConfig.AllowMethods
		MaxAge: sConfig.CorsConfig.MaxAge
	}

	app := fiber.New(sConfig)

	// adding middleware
	app.Use(cors.New(cConfig))
	app.Use(helmet.New())

	// TODO: add request timeout and other configs to the server
	err := app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {

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
			// time.Sleep(sConfig.WaitTimeBeforeKill)
			break bLoop
		}
	}

	// TODO: add grpc shutting down handling
	err = app.Shutdown()
	if err != nil {
		// logging in server terminaltion
	}

	logger.Warn("shutting down the server :(")
}
