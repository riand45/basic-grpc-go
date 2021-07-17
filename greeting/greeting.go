package greeting

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Receveid message body from client: %s", message.Body)
	return &Message{Body: "Hello from the server"}, nil
}