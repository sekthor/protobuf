package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sekthor/protobuf/api"
	"google.golang.org/grpc"
)

func main() {

	port := 5000

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	srv := Server{}
	api.RegisterPersonServiceServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
