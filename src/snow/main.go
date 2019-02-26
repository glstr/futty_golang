package main

import (
	"api"
	"confserver"
	"log"
	"net/http"
	_ "net/http/pprof"
	"snowplat"
	"utils"

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

	//middleware

	r.Use(gin.Logger())
	r.GET("/ping", snow)
	r.GET("/hello", helloWorld)
	r.StaticFile("/text", "./data/text.txt")

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

	//snowplat
	s := snowplat.NewSnowPlat(r)
	s.LoadRouter()
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

func snow(c *gin.Context) {
	c.JSON(200, gin.H{
		"host": "snow",
	})
}

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}
