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

type Server struct {
	protos.UnimplementedNameServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RegisterGrpc(srv *grpc.Server) error {
	protos.RegisterNameServer(srv, s)
	return nil
}

func (s *Server) RegisterHttp(ctx context.Context, mux *runtime.ServeMux,
	client *grpc.ClientConn) error {
	return protos.RegisterNameHandler(ctx, mux, client)
}

func (s *Server) GetName(ctx context.Context, in *protos.NameGetRequest) (*protos.NameGetResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (s *Server) SetName(ctx context.Context, in *protos.NameSetRequest) (*protos.NameSetResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalf("failed to initialize server: %s", err)
	}

	nh := NewServer()

	// register dual handlers
	srv.WithDualRegisterer(nh)

	err = srv.Run(context.Background(), HttpPort, GrpcPort)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
