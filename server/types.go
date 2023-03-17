package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GrpcRegisterer interface {
	RegisterGrpc(*grpc.Server) error
}

type HttpRegisterer interface {
	RegisterHttp(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
}

type Registry interface {
	GrpcRegisterer
	HttpRegisterer
}
