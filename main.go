package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化日志
	InitModule()

	//snow server
	r := gin.Default()
	s := service.NewSnowPlat(r)
	s.Load()
	r.Run(":8765")

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
