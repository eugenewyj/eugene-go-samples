package main

import (
	"fmt"
	"net"

	"github.com/eugenewyj/eugene-go-samples/grpc-sample/api"
	"github.com/ngaut/log"
	"google.golang.org/grpc"
)

// 启动一个gRPC服务，并等待连接
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("监听接口失败：%v", err)
	}

	//创建一个Server实例
	s := api.Server{}

	// 创建一个gRPC服务实例
	grpcServer := grpc.NewServer()

	// 绑定api服务到grpc服务
	api.RegisterPingServer(grpcServer, &s)

	// 启动服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动服务失败： %s", err)
	}
}
