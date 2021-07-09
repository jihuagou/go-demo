package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/tmaio/go-gin-example/models"
	"github.com/tmaio/go-gin-example/pkg/logging"
	"github.com/tmaio/go-gin-example/pkg/setting"
	"github.com/tmaio/go-gin-example/routers"
	"log"
	"syscall"
)

func main() {

	setting.Setup()
	models.Setup()

	//logging.Setup(logging.FileLog{})
	logging.Setup(logging.DbLog{})

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%v", err)
	}

}
