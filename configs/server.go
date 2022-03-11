package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rohanraj7316/middleware/configs"
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
func NewServerConfig() (*ServerConfigStruct, error) {
	kWaitTime, err := time.ParseDuration(GetValue("WAIT_TIME_BEFORE_KILL", "10s"))
	if err != nil {
		return nil, err
	}

	sConfig := configs.ServerDefault
	sConfig.AppName = GetValue("MODULE_NAME", "")

	return &ServerConfigStruct{
		CorsConfig: cors.Config{
			AllowOrigins: GetValue("ALLOW_CORS_ORIGIN", ""),
			AllowMethods: GetValue("ALLOW_CORS_METHODS", "POST"),
		},
		ModuleName:         GetValue("MODULE_NAME", ""),
		Port:               GetValue("PORT", "8081"),
		ProductName:        GetValue("PRODUCT_NAME", ""),
		ServerConfig:       sConfig,
		WaitTimeBeforeKill: kWaitTime,
		Version:            GetValue("VERSION", ""),
	}, nil
}
