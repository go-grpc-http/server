package version

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rohanraj7316/rsrc-bp-grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	protos.UnimplementedVersionServer
}

func New() *Handler {
	return &Handler{}
}

func (s *Handler) RegisterGrpc(srv *grpc.Server) error {
	protos.RegisterVersionServer(srv, s)
	return nil
}

func (s *Handler) RegisterHttp(ctx context.Context, mux *runtime.ServeMux, client *grpc.ClientConn) error {
	return protos.RegisterVersionHandler(ctx, mux, client)
}

func (s *Handler) Version(ctx context.Context, _ *protos.VersionRequest) (*protos.VersionResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
