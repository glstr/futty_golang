package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "snow",
		})
	})
	r.GET("hello", HelloWorld)
	r.Run()
}

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}
