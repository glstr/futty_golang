package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/httpserver/controller/handler"
)

func LoadRouter(e *gin.Engine) error {
	e.POST("/snow/cmdservice/cmd", handler.Cmd)

	e.POST("/snow/file/upload", handler.Upload)
	e.POST("/snow/file/download", handler.Download)

	e.GET("/snow/get_video", handler.GetVideo)
	e.GET("/snow/get_video_list", handler.GetVideoList)
	return nil
}
