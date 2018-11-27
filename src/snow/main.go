package main

import (
	"api"
	"io"
	"log"
	"os"
	"snowplat"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化日志
	initLog()
	//设置路由
	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/ping", snow)
	r.GET("/hello", helloWorld)
	r.StaticFile("/text", "./data/text.txt")

	f := api.NewFileOp(r)
	f.LoadRouter()

	//snowplat
	s := snowplat.NewSnowPlat(r)
	s.LoadRouter()
	r.Run(":8765")
}

func initLog() {
	gin.DisableConsoleColor()
	f, err := os.Create("log/snow.log")
	if err != nil {
		log.Printf("[init log fail, errMsg:%s]", err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
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
