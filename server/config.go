package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	httpPort string
	grpcPort string

	timeout time.Duration
	flags   []string

	reflection             bool
	registeredGrpcHandlers []GrpcRegisterer
	registeredHttpHandlers []HttpRegisterer

	options struct {
		grpc []*grpc.ServerOption
	}
}

var ConfigDefault = &Config{
	grpcPort:   ":8080",
	httpPort:   ":8090",
	reflection: true,
}

func configDefault(config ...Config) (*Config, error) {
	return ConfigDefault, nil
}

func (cfg *Config) httpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:        cfg.httpPort,
		Handler:     handler,
		ReadTimeout: cfg.timeout,
	}
}

func (cfg *Config) grpcServer() *grpc.Server {
	return grpc.NewServer()
}

func (cfg *Config) listener(ctx context.Context) (net.Listener, error) {
	return net.Listen("tcp", cfg.grpcPort)
}
