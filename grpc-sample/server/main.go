package main

import (
	"fmt"
	"net"

	"strings"

	"log"

	"net/http"

	"github.com/eugenewyj/go-samples/grpc-sample/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// private type for Context keys
type contextKey int

const clientIDKey contextKey = iota

func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Login" || headerName == "Password" {
		return headerName, true
	}
	return "", false
}

// 认证代理，检查登录的客户信息
func authenticateClient(ctx context.Context, server *api.Server) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientLogin := strings.Join(md["login"], "")
		clientPassword := strings.Join(md["password"], "")
		if clientLogin != "eugene" {
			return "", fmt.Errorf("不合法的用户： %s", clientLogin)
		}
		if clientPassword != "wang" {
			return "", fmt.Errorf("密码错误：%s", clientPassword)
		}

		log.Printf("合法客户： %s", clientLogin)
		return "42", nil
	}

	return "", fmt.Errorf("无认证信息")
}

// unaryInterceptor 采用当前context调用authenticateClient
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s, ok := info.Server.(*api.Server)
	if !(ok) {
		return nil, fmt.Errorf("unable to cast server")
	}
	clientID, err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, clientIDKey, clientID)
	return handler(ctx, req)
}

func startGRPCServer(address, certFile, keyFile string) error {
	// create a listener on TCP port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("监听接口失败： %v", err)
	}

	// 创建一个服务实例
	server := api.Server{}

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("不能加载TLS keys: %s", err)
	}

	// 创建一个带证书参数的gRPC options
	opts := []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor)}

	grpcServer := grpc.NewServer(opts...)

	api.RegisterPingServer(grpcServer, &server)

	// 启动服务
	log.Printf("启动HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("启动服务失败： %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress, certFile string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher))

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return fmt.Errorf("不能加载TLS cretifiate: %s", err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	err = api.RegisterPingHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("不能注册服务器： %s", err)
	}

	log.Printf("启动HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)

	return nil
}

// 启动一个gRPC服务，并等待连接
func main() {
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	certFile := "cert/server.crt"
	keyFile := "cert/server.key"

	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress, certFile, keyFile)
		if err != nil {
			log.Fatalf("启动gRPC服务失败： %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(restAddress, grpcAddress, certFile)
		if err != nil {
			log.Fatalf("启动REST服务失败： %s", err)
		}
	}()

	// 无限循环
	log.Printf("进入无限循环")
	select {}
}
