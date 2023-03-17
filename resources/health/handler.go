package health

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rohanraj7316/rsrc-bp-grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	protos.UnimplementedHealthServer
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) RegisterGrpc(srv *grpc.Server) error {
	protos.RegisterHealthServer(srv, s)
	return nil
}

func (s *Handler) RegisterHttp(ctx context.Context, mux *runtime.ServeMux, client *grpc.ClientConn) error {
	return protos.RegisterHealthHandler(ctx, mux, client)
}

func (s *Handler) Health(ctx context.Context, _ *protos.HealthRequest) (*protos.HealthResponse, error) {
	return nil, status.Error(codes.DeadlineExceeded, codes.DeadlineExceeded.String())
}
