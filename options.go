package server

import (
	"context"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func (s *server) httpServer(ctx context.Context, httpAddr string,
	handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    httpAddr,
		Handler: handler,
	}
}

func (s *server) grpcServer(ctx context.Context, opt ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opt...)
}

func (s *server) listener(ctx context.Context, grpcAddr string) (net.Listener, error) {
	return net.Listen("tcp", grpcAddr)
}

// WithDualRegisterer used for registering http and gRPC handlers
func (s *server) WithDualRegisterer(regs ...Registry) {
	// TODO: add flagging checks
	for index := range regs {
		s.registeredGrpcHandlers = append(s.registeredGrpcHandlers, regs[index])
		s.registeredHttpHandlers = append(s.registeredHttpHandlers, regs[index])
	}
}

// WithGrpcRegisterer used for registering only gRPC handlers
func (s *server) WithGrpcRegisterer(regs ...GrpcRegisterer) {
	// TODO: add flagging checks
	s.registeredGrpcHandlers = append(s.registeredGrpcHandlers, regs...)
}

// WithHttpRegisterer used for registering only http handlers
func (s *server) WithHttpRegisterer(regs ...HttpRegisterer) {
	// TODO: add flagging checks
	s.registeredHttpHandlers = append(s.registeredHttpHandlers, regs...)
}

// WithUnaryServerInterceptor used for registering only [UnaryServerInterceptor](https://grpc.io/blog/grpc-web-interceptor/)
func (s *server) WithUnaryServerInterceptor(regs ...HttpRegisterer) {
	// TODO: add flagging checks
	s.registeredHttpHandlers = append(s.registeredHttpHandlers, regs...)
}
