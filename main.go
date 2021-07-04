package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/glstr/futty_golang/httpserver"
	"github.com/glstr/futty_golang/logger"
)

func main() {
	//初始化日志
	InitModule()

	logger.Notice("start")
	err := httpserver.StartHttpServer()
	if err != nil {
		panic(err)
	}

	//snow server
	//r := gin.Default()
	//s := service.NewSnowPlat(r)
	//s.Load()
	//r.Run(":8765")

}

func InitModule() {
	logPath := "log/snow.log"
	option := new(logger.LogOption)
	option.LogPath = logPath
	err := logger.InitLogger(option)
	if err != nil {
		panic(err)
	}

	go func() {
		err := http.ListenAndServe(":8764", nil)
		if err != nil {
			panic(err)
		}
	}()
}
