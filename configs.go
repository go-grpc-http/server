package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// ServerConfig is a struct which holds server config
type config struct {
	// TODO: can discuss out abt the naming convention more
	// represents name of the server or if you want it to be at service level
	Name string

	// When set true, it will enable reflectionFlag
	//
	// Default: false
	ReflectionFlag bool

	// When not set, it will log it error using default error handler
	//
	// Default: DefaultErrorHandler
	ErrorHandler ErrorHandler

	httpPort               string
	gRpcPort               string
	requestTimeout         time.Duration
	registeredGrpcHandlers []GrpcRegisterer
	registeredHttpHandlers []HttpRegisterer
	serverOpts             []grpc.ServerOption
}

func configDefault() (*config, error) {
	cfg := &config{}
	cfg.ReflectionFlag = true
	cfg.ErrorHandler = func(ctx context.Context, err error) error { return nil }

	return cfg, nil
}

// WithName used to name the server
func WithName(name string) Option {
	return func(c *config) error {
		c.Name = name

		return nil
	}
}

// WithDualRegisterer used for registering http and gRPC handlers
func WithDualRegisterer(regs ...Registry) Option {
	return func(c *config) error {
		// TODO: add flagging checks
		for index := range regs {
			c.registeredGrpcHandlers = append(c.registeredGrpcHandlers, regs[index])
			c.registeredHttpHandlers = append(c.registeredHttpHandlers, regs[index])
		}

		return nil
	}
}

// WithGrpcRegisterer used for registering only gRPC handlers
func WithGrpcRegisterer(regs ...GrpcRegisterer) Option {
	return func(c *config) error {
		// TODO: add flagging checks
		c.registeredGrpcHandlers = append(c.registeredGrpcHandlers, regs...)

		return nil
	}
}

// WithUnaryServerInterceptor used for registering only [UnaryServerInterceptor](https://grpc.io/blog/grpc-web-interceptor/)
func WithUnaryServerInterceptor(intrcptrs ...grpc.UnaryServerInterceptor) Option {
	return func(c *config) error {
		c.serverOpts = append(c.serverOpts, grpc.ChainUnaryInterceptor(intrcptrs...))

		return nil
	}
}

// WithErrHandler used to assign ErrHandler
func (c *config) WithErrHandler(fn ErrorHandler) Option {
	return func(c *config) error {
		c.ErrorHandler = fn

		return nil
	}
}

// WithHttpPort used to assign http port
func WithHttpPort(port int) Option {
	return func(c *config) error {
		c.httpPort = fmt.Sprintf(":%d", port)

		return nil
	}
}

// WithGrpcPort used to assign grpc port
func WithGrpcPort(port int) Option {
	return func(c *config) error {
		c.gRpcPort = fmt.Sprintf(":%d", port)

		return nil
	}
}

// WithRequestTimeout used to assign request timeout
func WithRequestTimeout(t time.Duration) Option {
	return func(c *config) error {
		c.requestTimeout = t

		return nil
	}
}
