package server

import (
	"net/http"

	"google.golang.org/grpc"
)

type server struct {
	cfg        *config
	httpServer *http.Server
	gRpcServer *grpc.Server
}

// New used to initialize new server.
func New(opts ...Option) (*server, error) {
	cfg, err := configDefault()
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err = opt(cfg)
		if err != nil {
			return nil, err
		}
	}

	return &server{
		cfg: cfg,
	}, nil
}
