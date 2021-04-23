package main

import (
	"github.com/glstr/futty_golang/api"
	"github.com/glstr/futty_golang/confserver"
	"github.com/glstr/futty_golang/service"
	"github.com/glstr/futty_golang/utils"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化日志
	initLog()
	go func() {
		log.Println(http.ListenAndServe("localhost:8764", nil))
	}()
	//设置路由
	r := gin.Default()

	//init module

	//middleware

	//file server
	f := api.NewFileOp(r)
	f.LoadRouter()

	//config server
	confFile := "./conf/confserver.conf"
	confServer := confserver.NewConfServer(r)
	err := confServer.ServiceInit(confFile)
	if err != nil {
		log.Printf("[err_msg:%s]", err.Error())
		return
	}
	confServer.LoadRouter()

	s := service.NewSnowPlat(r)
	s.LoadRouter()

	//start service
	r.Run(":8765")
}

func initLog() {
	utils.LogInit("log/snow.log")
	//gin.DisableConsoleColor()
	//f, err := os.Create("log/snow.log")
	//if err != nil {
	//	log.Printf("[init log fail, errMsg:%s]", err.Error())
	//}
	//gin.DefaultWriter = io.MultiWriter(f)
	//log.SetOutput(gin.DefaultWriter)
}
