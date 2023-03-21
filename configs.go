package server

import (
	"errors"
	"time"
)

var (
	ErrServerConfigNameMissing           = errors.New("")
	ErrServerConfigRequestTimeoutMissing = errors.New("")
)

// ServerConfig is a struct which holds server config
type ServerConfig struct {
	// TODO: can discuss out abt the naming convention more
	// represents name of the server or if you want it to be at service level
	Name string

	// When not set, it will use default timeout
	//
	// Default: 1 sec
	RequestTimeout time.Duration

	// When set true, it will enable reflectionFlag
	//
	// Default: false
	ReflectionFlag bool
}

// ServerConfigDefault is set of default sever configs
var ServerConfigDefault = ServerConfig{
	RequestTimeout: 1 * time.Second,
	ReflectionFlag: true,
}

func serverConfigDefault(config ...ServerConfig) (cfg ServerConfig, err error) {
	return ServerConfigDefault, err
}
