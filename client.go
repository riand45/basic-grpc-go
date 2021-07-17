package main

import (
	"context"
	"github.com/riand45/basic-grpc-go/greeting"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}
	defer conn.Close()

	g := greeting.NewGreetingMessageClient(conn)

	message := greeting.Message{
		Body: "Hello from client",
	}

	response, err := g.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}
