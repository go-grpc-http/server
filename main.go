package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rohanraj7316/rsrc-bp-grpc/trade"
	"google.golang.org/grpc"
)

type server struct {
	trade.UnimplementedTradingMSServer
}

func NewServer() *server {
	return &server{}
}

// @reference - https://github.com/grpc-ecosystem/grpc-gateway/issues/2039#issuecomment-799560929
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Greeter service to the server
	trade.RegisterTradingMSServer(s, &server{})

	// Serve gRPC Server
	log.Println("serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	gmux := runtime.NewServeMux()

	err = trade.RegisterTradingMSHandler(context.Background(), gmux, conn)
	if err != nil {
		log.Fatalln("failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gmux,
	}

	log.Println("serving grpc-gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
