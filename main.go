package main

import (
	"fmt"
	"log"
	_ "net/http"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/hucongyang/go-demo/conf"
	"github.com/hucongyang/go-demo/routers"
)

func main() {
	// 新版本服务：优雅的重启服务
	conf := conf.Config()
	endless.DefaultReadTimeOut = conf.Server.ReadTimeout
	endless.DefaultWriteTimeOut = conf.Server.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", conf.Server.HttpPort)

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
	//conf := conf.Config()
	//server := &http.Server{
	//	Addr: fmt.Sprintf(":%d", conf.Server.HttpPort),
	//	Handler: router,
	//	ReadTimeout: conf.Server.ReadTimeout,
	//	WriteTimeout: conf.Server.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()
}
