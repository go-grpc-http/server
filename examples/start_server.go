package main

import (
	"context"
	"log"

	"github.com/go-grpc-http/server"
	"github.com/go-grpc-http/server/examples/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	HttpPort = 8080
	GrpcPort = 8090
)

type Handler struct {
	protos.UnimplementedNameServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (s *Handler) RegisterGrpc(srv *grpc.Server) error {
	protos.RegisterNameServer(srv, s)
	return nil
}

func (s *Handler) RegisterHttp(ctx context.Context, mux *runtime.ServeMux,
	client *grpc.ClientConn) error {
	return protos.RegisterNameHandler(ctx, mux, client)
}

func (s *Handler) GetName(ctx context.Context, in *protos.NameGetRequest) (*protos.NameGetResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (s *Handler) SetName(ctx context.Context, in *protos.NameSetRequest) (*protos.NameSetResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func main() {
	nh := NewHandler()

	srv, err := server.New(
		server.WithHttpPort(HttpPort),
		server.WithGrpcPort(GrpcPort),
		server.WithDualRegisterer(nh),
	)
	if err != nil {
		log.Fatalf("failed to initialize server: %s", err)
	}

	err = srv.Run(context.Background())
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
