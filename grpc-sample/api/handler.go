package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server 代表gRPC服务端
type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("收到消息： %s", in.Greeting)
	return &PingMessage{Greeting: "OK"}, nil
}
