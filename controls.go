package server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/go-grpc-http/server/resources"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func (s *server) start(ctx context.Context, httpAddr, grpcAddr string) error {

	lis, err := s.listener(ctx, grpcAddr)
	if err != nil {
		return err
	}

	s.gRpcServer = s.initGrpcServer(ctx, s.cfg.serverOpts...)

	// initialize health and version routes
	ih := resources.New()

	s.cfg.registeredGrpcHandlers = append(s.cfg.registeredGrpcHandlers, []GrpcRegisterer{ih}...)
	s.cfg.registeredHttpHandlers = append(s.cfg.registeredHttpHandlers, []HttpRegisterer{ih}...)

	for _, handler := range s.cfg.registeredGrpcHandlers {
		err = handler.RegisterGrpc(s.gRpcServer)
		if err != nil {
			return err
		}
	}

	if s.cfg.ReflectionFlag {
		reflection.Register(s.gRpcServer)
	}

	if s.cfg.httpPort != "" {
		eg := errgroup.Group{}

		eg.Go(func() error {
			conn, err := grpc.DialContext(
				ctx,
				fmt.Sprintf("0.0.0.0%s", grpcAddr),
				grpc.WithBlock(),
				grpc.WithTransportCredentials(
					insecure.NewCredentials(),
				),
			)
			if err != nil {
				return err
			}

			mux := runtime.NewServeMux()

			for _, handler := range s.cfg.registeredHttpHandlers {
				err = handler.RegisterHttp(ctx, mux, conn)
				if err != nil {
					return err
				}
			}

			s.httpServer = s.initHttpServer(ctx, httpAddr, mux)

			log.Info().Msgf("serving grpc-gateway on http://0.0.0.0%s", httpAddr)
			return s.httpServer.ListenAndServe()
		})
	}

	log.Info().Msgf("serving gRPC on 0.0.0.0%s", grpcAddr)
	return s.gRpcServer.Serve(lis)
}

// TODO: need to implement
func (s *server) stop(ctx context.Context) error {
	return nil
}

// TODO: need to implement
func (s *server) cleanup(ctx context.Context) error {
	return nil
}

// Run use to start running your server
func (s *server) Run(ctx context.Context) error {
	return s.start(ctx, s.cfg.httpPort, s.cfg.gRpcPort)
}

// TODO: need to improve the implementation
// Stop use to stop your server gracefully
func (s *server) Stop(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	s.gRpcServer.GracefulStop()

	return nil
}
