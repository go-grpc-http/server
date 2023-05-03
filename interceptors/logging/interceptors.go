package logging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog"
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

		go func(info *grpc.UnaryServerInfo, req, resp interface{}, start time.Time, logger *zerolog.Event, err error) {
			if err != nil {
				logger = log.Error().Err(err)
			}

			statusCode := codes.Unknown
			if st, ok := status.FromError(err); ok {
				statusCode = st.Code()
			}

			latency := time.Since(start)

			if req != nil {
				reqByte, err := json.Marshal(req)
				if err != nil {
					logger.Err(err)
				}

				logger.Str("reqBody", string(reqByte))
			}

			if resp != nil {
				respByte, err := json.Marshal(resp)
				if err != nil {
					logger.Err(err)
				}

				logger.Str("resBody", string(respByte))
			}

			logger.Str("protocol", "grpc").
				Str("method", info.FullMethod).
				Int("statusCode", int(statusCode)).
				Str("status", statusCode.String()).
				Dur("latency", latency).
				Msgf("[REQ-RES-LOG] %d %s %s", statusCode, info.FullMethod, latency) // [REQ-RES-LOG] 0 /health 2.3sec
		}(info, req, resp, start, logger, err)

		return resp, err
	}
}
