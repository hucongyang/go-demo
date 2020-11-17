package main

import (
	"fmt"
	"github.com/hucongyang/go-demo/pkg/gredis"
	"log"
	_ "net/http"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/hucongyang/go-demo/conf"
	"github.com/hucongyang/go-demo/routers"
)

func main() {
	// 新版本服务：优雅的重启服务
	_ = gredis.Setup()
	config := conf.Config()
	endless.DefaultReadTimeOut = config.Server.ReadTimeout
	endless.DefaultWriteTimeOut = config.Server.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", config.Server.HttpPort)

	router := routers.InitRouter()
	server := endless.NewServer(endPoint, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	// 老版本服务
	//router := routers.InitRouter()
	//config := config.Config()
	//server := &http.Server{
	//	Addr: fmt.Sprintf(":%d", config.Server.HttpPort),
	//	Handler: router,
	//	ReadTimeout: config.Server.ReadTimeout,
	//	WriteTimeout: config.Server.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()
}
