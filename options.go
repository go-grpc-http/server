package server

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GrpcRegisterer interface {
	RegisterGrpc(*grpc.Server) error
}

type HttpRegisterer interface {
	RegisterHttp(context.Context, *runtime.ServeMux,
		*grpc.ClientConn) error
}

type Registry interface {
	GrpcRegisterer
	HttpRegisterer
}

type Option func(*config) error

type ErrorHandler func(ctx context.Context, err error) error

func (s *server) initHttpServer(ctx context.Context, httpAddr string,
	handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    httpAddr,
		Handler: handler,
	}
}

func (s *server) initGrpcServer(ctx context.Context, opts ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}

func (s *server) listener(ctx context.Context, grpcAddr string) (net.Listener, error) {
	return net.Listen("tcp", grpcAddr)
}
