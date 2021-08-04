package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/model/service"
)

type GetVideoRequest struct {
	VideoId int64 `form:"video_id"`
}

type GetVideoResponse struct {
	middleware.CommonResponse
	Data *service.VideoInfo `json:"data"`
}

func GetVideo(c *gin.Context) {
	ctx := context.GetContext()
	var res GetVideoResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	ctx.LogBuffer.WriteLog("method[GetVideo] ")
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	//get param
	var req GetVideoRequest
	err = middleware.GetJsonParam(c, &req)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("video_id[%d]", req.VideoId)

	//call service
	ser := service.GetVideoService()
	var videoInfo *service.VideoInfo
	videoInfo, err = ser.GetVideo(req.VideoId)
	if err != nil {
		ctx.LogBuffer.WriteLog("error_msg[%s]", err.Error())
		return
	}

	//make response
	res.Data = videoInfo
}

type GetVideoListResponse struct {
	middleware.CommonResponse
	VideoIdList []*service.VideoInfo `json:"video_id_list"`
}

func GetVideoList(c *gin.Context) {
	ctx := context.GetContext()
	var res GetVideoListResponse
	res.CommonResponse.RequestId = ctx.Logid
	ctx.LogBuffer.WriteLog("method[GetVideoList] ")
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		logger.Notice(ctx.LogBuffer.String())
		c.JSON(200, res)
		context.PutContext(ctx)
	}()

	//call service
	ser := service.GetVideoService()
	var list []*service.VideoInfo
	list, err = ser.GetVideoList()
	if err != nil {
		ctx.LogBuffer.WriteLog("error_msg[%s]", err.Error())
		return
	}

	ctx.LogBuffer.WriteLog("res[%v]", list)
	res.VideoIdList = list
}
