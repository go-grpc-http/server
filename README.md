# Server

Server is an inspired [google.api.http](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto#L46). Designed to ease implementation and pace up the development. internally it uses [protoc](https://github.com/protocolbuffers/protobuf) which reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC.

This helps you provide your APIs in both gRPC and RESTful style at the same time.

## Kickstart

```
import (
	"context"
	"log"

	"github.com/go-grpc-http/server"
)

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatalf("failed to initialize server: %s", err)
	}

	err = srv.Run(context.Background(), HttpPort, GrpcPort)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
```

## Usage

### 1. Define your protobuf

```
syntax = "proto3";

option go_package="/protos";

import "google/api/annotations.proto";

// define empty message
message NameGetRequest {}

message NameGetResponse {
    int32 statusCode = 1;
    string status = 2;
    string message = 3;
}

// define empty message
message NameSetRequest {}

message NameSetResponse {
    int32 statusCode = 1;
    string status = 2;
    string message = 3;
}

service Name {
    rpc SetName(NameSetRequest) returns(NameSetResponse) {
        option (google.api.http) = {
            post: "/"
            body: "*"
        };
    }
    rpc GetName(NameGetRequest) returns(NameGetResponse) {
        option (google.api.http) = {
            get: "/"
            body: "*"
        };
    }
}
```

