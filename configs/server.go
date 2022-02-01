package configs

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// ServerConfig config.
type ServerConfigStruct struct {
	Port               string
	CorsConfig         cors.Config
	DisableMaintenance string
	ProductName        string
	ModuleName         string
	ServerConfig       fiber.Config
	WaitTimeBeforeKill time.Duration
	Version            string
}

// return server config.
func NewServerConfig(cancel context.CancelFunc) *ServerConfigStruct {
	rTimeout, err := time.ParseDuration(GetValue("READ_TIME_OUT", ""))
	if err != nil {
		// TODO: log/kill the server
		cancel()
	}

	kWaitTime, err := time.ParseDuration(GetValue("WAIT_TIME_BEFORE_KILL", "10s"))
	if err != nil {
		// TODO: log the server
	}

	// mTime, err := strconv.Atoi(GetValue("REQUEST_MAX_AGE", "10"))
	// if err != nil {
	// 	cancel()
	// }

	bLimit, err := strconv.Atoi(GetValue("BODY_LIMIT", ""))
	if err != nil {
		cancel()
	}

	return &ServerConfigStruct{
		CorsConfig: cors.Config{
			AllowOrigins: GetValue("ALLOW_CORS_ORIGIN", ""),
			AllowMethods: GetValue("ALLOW_CORS_METHODS", "POST"),
		},
		ModuleName:  GetValue("MODULE_NAME", ""),
		Port:        GetValue("PORT", "8081"),
		ProductName: GetValue("PRODUCT_NAME", ""),
		ServerConfig: fiber.Config{
			BodyLimit:   bLimit,
			ReadTimeout: rTimeout,
		},
		WaitTimeBeforeKill: kWaitTime,
		Version:            GetValue("VERSION", ""),
	}
}
