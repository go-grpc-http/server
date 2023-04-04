package health

import (
	"context"
	"net/http"

	"github.com/go-grpc-http/server/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
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
	return &protos.HealthResponse{
		StatusCode: http.StatusOK,
		Message:    "Health Successful",
	}, nil
}
