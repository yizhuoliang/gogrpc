package chat

import (
	"log"
	
	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) SayHello(ctx context.Context, message *ChatMessage) (*ChatMessage, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &ChatMessage{Body: "Hello From the Server"}, nil
}