package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/model/service"
)

type CommonResponse struct {
	ErrorCode int32  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	RequestId int64  `json:"request_id"`
}

type GetVideoRequest struct {
	VideoId int64 `form:"video_id"`
}

type GetVideoResponse struct {
	CommonResponse
	Data *service.VideoInfo `json:"data"`
}

func GetVideo(c *gin.Context) {
	ctx := context.GetContext()
	var res GetVideoResponse
	res.CommonResponse.RequestId = ctx.Logid
	ctx.LogBuffer.WriteLog("method[GetVideo] ")
	defer func() {
		logger.Notice(ctx.LogBuffer.String())
		c.JSON(200, res)
		context.PutContext(ctx)
	}()

	//get param
	var req GetVideoRequest
	err := c.ShouldBind(&req)
	if err != nil {
		ctx.LogBuffer.WriteLog("error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("video_id[%d]", req.VideoId)

	//check param

	//call service
	ser := service.GetVideoService()
	vedioInfo, err := ser.GetVideo(req.VideoId)
	if err != nil {
		ctx.LogBuffer.WriteLog("error_msg[%s]", err.Error())
		return
	}

	//make response
	res.Data = vedioInfo
}

type GetVideoListResponse struct {
	CommonResponse
	VideoIdList []*service.VideoInfo `json:"video_id_list"`
}

func GetVideoList(c *gin.Context) {
	ctx := context.GetContext()
	var res GetVideoListResponse
	res.CommonResponse.RequestId = ctx.Logid
	ctx.LogBuffer.WriteLog("method[GetVideoList] ")
	defer func() {
		logger.Notice(ctx.LogBuffer.String())
		c.JSON(200, res)
		context.PutContext(ctx)
	}()

	//call service
	ser := service.GetVideoService()
	list, err := ser.GetVideoList()
	if err != nil {
		ctx.LogBuffer.WriteLog("error_msg[%s]", err.Error())
		return
	}

	res.VideoIdList = list
}
