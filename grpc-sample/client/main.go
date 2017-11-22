package main

import (
	"google.golang.org/grpc"

	"log"

	"github.com/eugenewyj/eugene-go-samples/grpc-sample/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

// Authentication holds the login/password
type Authentication struct {
	Login    string
	Password string
}

// GetRequestMetadata 获取当前请求元数据
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity 标识是否认证需要安全传输
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("不能加载tls cert: %s", err)
	}

	auth := Authentication{
		Login:    "eugene",
		Password: "wang",
	}

	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
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
