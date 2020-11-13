package main

import (
	"fmt"
	"github.com/hucongyang/go-demo/conf"
	"github.com/hucongyang/go-demo/routers"
	"net/http"
)

func main()  {
	router := routers.InitRouter()
	conf := conf.Config()
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", conf.Server.HttpPort),
		Handler: router,
		ReadTimeout: conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
