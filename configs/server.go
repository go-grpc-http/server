package configs

import (
	"context"
	"strconv"
	"time"
)

type CorsConfigstruct struct {
	// AllowOrigin defines a list of origins that may access the resource.
	//
	// Optional. Default value "*"
	AllowOrigins string

	// AllowMethods defines a list methods allowed when accessing the resource.
	// This is used in response to a preflight request.
	//
	// Optional. Default value "GET,POST,HEAD,PUT,DELETE,PATCH"
	AllowMethods string

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached.
	//
	// Optional. Default value 0.
	MaxAge int
}

// ServerConfig config.
type ServerConfigStruct struct {
	Port               string
	GrpcPort           string
	ReadTimeout        time.Duration
	ReadHeaderTimeout  time.Duration
	DisableMaintenance string
	ProductName        string
	ModuleName         string
	WaitTimeBeforeKill time.Duration
	CorsConfig         CorsConfigstruct
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

	mTime, err := strconv.Atoi(GetValue("REQUEST_MAX_AGE", "10"))
	if err != nil {
		cancel()
	}

	return &ServerConfigStruct{
		Port:               GetValue("PORT", "8081"),
		GrpcPort:           GetValue("GRPC_PORT", "8080"),
		ReadTimeout:        rTimeout,
		ModuleName:         GetValue("MODULE_NAME", ""),
		ProductName:        GetValue("PRODUCT_NAME", ""),
		WaitTimeBeforeKill: kWaitTime,
		CorsConfig: CorsConfigstruct{
			AllowOrigins: GetValue("ALLOW_CORS_ORIGIN", ""),
			AllowMethods: GetGetValue("ALLOW_CORS_METHODS", "POST")
			MaxAge: mTime,
		},
	}
}
