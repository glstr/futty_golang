package main

import (
	"api"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化日志
	initLog()
	//设置路由
	r := gin.Default()
	r.GET("/ping", snow)
	r.GET("/hello", helloWorld)
	r.StaticFile("/text", "./data/text.txt")

	f := api.NewFileOp(r)
	f.LoadRouter()

	r.Run(":8765")
}

func initLog() {
	gin.DisableConsoleColor()
	f, err := os.Create("log/snow.log")
	if err != nil {
		log.Printf("[init log fail, errMsg:%s]", err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f)
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
