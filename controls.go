package server

import (
	"context"
	"fmt"
	"log"

	"github.com/go-grpc-http/server/resources/health"
	"github.com/go-grpc-http/server/resources/version"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func (s *server) init(ctx context.Context) error {
	return nil
}

func (s *server) start(ctx context.Context, httpAddr, grpcAddr string) error {

	lis, err := s.listener(ctx, grpcAddr)
	if err != nil {
		return err
	}

	/*
		TODO: define way to add unary interceptor
	*/

	s.gServer = s.grpcServer(ctx)

	// initialize health and version routes
	hh := health.New()
	vh := version.New()
	s.registeredGrpcHandlers = append(s.registeredGrpcHandlers, []GrpcRegisterer{hh, vh}...)
	s.registeredHttpHandlers = append(s.registeredHttpHandlers, []HttpRegisterer{hh, vh}...)

	for _, handler := range s.registeredGrpcHandlers {
		err = handler.RegisterGrpc(s.gServer)
		if err != nil {
			log.Fatal(err)
		}
	}

	if s.cfg.ReflectionFlag {
		reflection.Register(s.gServer)
	}

	log.Printf("serving gRPC on 0.0.0.0%s", grpcAddr)
	go func() {
		s.gServer.Serve(lis)
	}()

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

	gwMux := runtime.NewServeMux()

	for _, handler := range s.registeredHttpHandlers {
		err = handler.RegisterHttp(ctx, gwMux, conn)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("serving grpc-gateway on http://0.0.0.0%s", httpAddr)
	s.hServer = s.httpServer(ctx, httpAddr, gwMux)

	return s.hServer.ListenAndServe()
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
func (s *server) Run(ctx context.Context, httpPort, grpcPort int) error {
	httpAddr := fmt.Sprintf(":%d", httpPort)
	grpcAddr := fmt.Sprintf(":%d", grpcPort)

	// TODO: add listener handling
	return s.start(ctx, httpAddr, grpcAddr)
}

// TODO: need to improve the implementation
// Stop use to stop your server gracefully
func (s *server) Stop(ctx context.Context) error {
	err := s.hServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	s.gServer.GracefulStop()

	return nil
}
