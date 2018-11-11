package main

import (
	"log"
	"net/http"

	"os"

	"os/signal"
	"syscall"

	"github.com/eugenewyj/go-samples/advent/handler"
	"github.com/eugenewyj/go-samples/advent/version"
	"golang.org/x/net/context"
)

func main() {
	log.Printf("启动服务......\n提交注释: %s, 编译时间：%s, 版本号：%s", version.Commit, version.BuildTime, version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("未指定端口")
	}

	r := handler.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Print("服务已启动，开始监听。")

	killSignal := <-interrupt
	switch killSignal {
	case os.Kill:
		log.Print("Got SIGKILL...")
	case os.Interrupt:
		log.Print("Got SIGNT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("服务器正在关闭...")
	srv.Shutdown(context.Background())
	log.Print("服务器已经关闭。")
}
