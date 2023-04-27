package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor used for logging request and response
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger, start := log.Info(), time.Now()

		resp, err = handler(ctx, req)
		if err != nil {
			logger = log.Error()
		}

		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		latency := time.Since(start)
		logger.Str("protocol", "grpc").
			Str("method", info.FullMethod).
			Int("statusCode", int(statusCode)).
			Str("status", statusCode.String()).
			Dur("latency", latency).
			Interface("reqBody", req).
			Interface("resBody", resp).
			Msg(fmt.Sprintf("[REQ-RES-LOG] %d %s %s", statusCode, info.FullMethod, latency)) // [REQ-RES-LOG] 0 /health 2.3sec

		return resp, err
	}
}
