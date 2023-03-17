package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	ENV_HTTP_PORT       = "HTTP_PORT"
	ENV_GRPC_PORT       = "GRPC_PORT"
	ENV_REQUEST_TIMEOUT = "REQUEST_TIMEOUT"
)

type Config struct {
	httpPort string
	grpcPort string

	timeout time.Duration
}

var ConfigDefault = &Config{}

func configDefault(config ...Config) (*Config, error) {
	httpPort := os.Getenv(ENV_HTTP_PORT)
}

func (cfg *Config) httpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:        cfg.httpPort,
		Handler:     handler,
		ReadTimeout: cfg.timeout,
	}
}

func (cfg *Config) grpcServer(opt ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opt...)
}

func (cfg *Config) listener(ctx context.Context) (net.Listener, error) {
	return net.Listen("tcp", cfg.grpcPort)
}
