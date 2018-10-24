package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initLog()
	r := gin.Default()
	r.GET("/ping", snow)
	r.GET("/hello", helloWorld)
	r.POST("/upload", upload)
	r.Run()
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

func upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	defaultPath := "./data/"
	dst := defaultPath + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Printf("[save file fail, errMsg:%s]", err.Error())
	}
	c.String(http.StatusOK, fmt.Sprintf("%s upload", file.Filename))
}
