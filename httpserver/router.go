package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/httpserver/controller/handler"
)

func LoadRouter(e *gin.Engine) error {
	//e.GET("/snow/get_pic", handler.GetPic)

	e.GET("/snow/get_video", handler.GetVideo)
	e.GET("/snow/get_video_list", handler.GetVideoList)
	return nil
}
