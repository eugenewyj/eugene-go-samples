package main

import (
	"google.golang.org/grpc"

	"log"

	"github.com/eugenewyj/eugene-go-samples/grpc-sample/api"
	"golang.org/x/net/context"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接失败： %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Hello"})
	if err != nil {
		log.Fatalf("调用SayHello出错： %s", err)
	}
	log.Printf("服务器响应： %s", response.Greeting)
}
