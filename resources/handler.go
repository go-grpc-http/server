package resources

import (
	"context"
	"net/http"

	"github.com/go-grpc-http/server/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	protos.UnimplementedInternalRoutesServer
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) RegisterGrpc(srv *grpc.Server) error {
	protos.RegisterInternalRoutesServer(srv, s)
	return nil
}

func (s *Handler) RegisterHttp(ctx context.Context, mux *runtime.ServeMux, client *grpc.ClientConn) error {
	return protos.RegisterInternalRoutesHandler(ctx, mux, client)
}

func (s *Handler) Health(ctx context.Context, _ *emptypb.Empty) (*protos.HealthResponse, error) {
	return &protos.HealthResponse{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Message:    "Health Successful",
	}, nil
}

func (s *Handler) Version(context.Context, *emptypb.Empty) (*protos.VersionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "coming soon !!")
}
