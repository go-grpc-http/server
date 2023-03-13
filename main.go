package main

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rohanraj7316/rsrc-bp-grpc/server"
	"github.com/rohanraj7316/rsrc-bp-grpc/trade"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	trade.UnimplementedTradingMSServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RegisterGrpc(srv *grpc.Server) error {
	trade.RegisterTradingMSServer(srv, NewServer())
	return nil
}

func (s *Server) RegisterHttp(ctx context.Context, mux *runtime.ServeMux, client *grpc.ClientConn) error {
	return trade.RegisterTradingMSHandler(ctx, mux, client)
}

func (s *Server) Health(ctx context.Context, _ *trade.HealthRequest) (*trade.HealthResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

// @reference - https://github.com/grpc-ecosystem/grpc-gateway/issues/2039#issuecomment-799560929
func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalln(err)
	}

	srv.WithDualRegisterer(NewServer())

	srv.Run(context.Background())
}
