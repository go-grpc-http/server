package main

import (
	"context"
	"go-grpc-http/resources/health"
	"go-grpc-http/resources/version"
	"go-grpc-http/server"
	"log"
)

// @reference - https://github.com/grpc-ecosystem/grpc-gateway/issues/2039#issuecomment-799560929
func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalln(err)
	}

	hHandler := health.New()
	vHandler := version.New()

	resources := []server.Registry{hHandler, vHandler}
	srv.WithDualRegisterer(resources...)

	srv.Run(context.Background())
}
