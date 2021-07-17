package main

import (
	"github.com/riand45/basic-grpc-go/greeting"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := greeting.Server{}

	grpcServer := grpc.NewServer()

	greeting.RegisterGreetingMessageServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
