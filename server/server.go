package server

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	cfg *Config

	grpcServer *grpc.Server
}

func New(config ...Config) (*server, error) {
	cfg, err := configDefault(config...)
	if err != nil {

	}

	// Create a gRPC server object
	return &server{
		cfg: cfg,
	}, nil
}

func (s *server) WithDualRegisterer(regs ...Registry) {
	for index := range regs {
		s.cfg.registeredGrpcHandlers = append(s.cfg.registeredGrpcHandlers, regs[index])
		s.cfg.registeredHttpHandlers = append(s.cfg.registeredHttpHandlers, regs[index])
	}
}

func (s *server) Run(ctx context.Context) {
	lis, err := s.cfg.listener(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	s.grpcServer = s.cfg.grpcServer()

	for _, handler := range s.cfg.registeredGrpcHandlers {
		err = handler.RegisterGrpc(s.grpcServer)
		if err != nil {
			log.Fatal(err)
		}
	}

	if s.cfg.reflection {
		reflection.Register(s.grpcServer)
	}

	log.Printf("serving gRPC on 0.0.0.0%s", s.cfg.grpcPort)
	go func() {
		log.Fatalln(s.grpcServer.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("0.0.0.0%s", s.cfg.grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	gwMux := runtime.NewServeMux()

	for _, handler := range s.cfg.registeredHttpHandlers {
		err = handler.RegisterHttp(ctx, gwMux, conn)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("serving grpc-gateway on http://0.0.0.0:%s", s.cfg.httpPort)
	gwServer := s.cfg.httpServer(gwMux)
	log.Fatalln(gwServer.ListenAndServe())
}
