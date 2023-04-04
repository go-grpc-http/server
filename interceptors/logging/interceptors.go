package logging

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		// start := time.Now()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
