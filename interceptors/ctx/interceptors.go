package ctx

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// TODO: still need to add logic for relay-back & pass-on headers
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if mds, ok := metadata.FromIncomingContext(ctx); ok {
			md := metadata.Join(mds)
			grpc.SendHeader(ctx, md)
		}

		return handler(ctx, req)
	}
}
