package main

import (
	"context"
	"log"

	"github.com/rohanraj7316/rsrc-bp-grpc/resources/health"
	"github.com/rohanraj7316/rsrc-bp-grpc/resources/version"
	"github.com/rohanraj7316/rsrc-bp-grpc/server"
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
