package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartHttpServer() error {
	e := gin.Default()
	
	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true  // 允许所有来源，生产环境建议指定具体域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	e.Use(cors.New(config))
	
	err := LoadRouter(e)
	if err != nil {
		return err
	}
	e.Run(":8882")
	return nil
}
