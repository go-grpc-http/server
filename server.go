package server

import (
	"net/http"

	"google.golang.org/grpc"
)

type server struct {
	cfg ServerConfig

	hServer *http.Server
	gServer *grpc.Server

	registeredGrpcHandlers []GrpcRegisterer
	registeredHttpHandlers []HttpRegisterer
}

// New used to initialize config for server.
func New(config ...ServerConfig) (*server, error) {
	cfg, err := serverConfigDefault(config...)
	if err != nil {
		return nil, err
	}

	// return struct that been used internally
	return &server{
		cfg: cfg,
	}, nil
}
