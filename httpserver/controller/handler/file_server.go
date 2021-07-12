package handler

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
)

const DefaultDir = "./data"

type UploadRequest struct {
	Container string `json:"container"`
	Key       string `json:"key"`
}

type UploadResponse struct {
	middleware.CommonResponse
	FileName string `json:"file_name"`
}

func Upload(c *gin.Context) {
	ctx := context.GetContext()
	var res UploadResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	var req UploadRequest
	if err = middleware.GetJsonParam(c, &req); err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}

	container := c.PostForm("container")
	key := c.PostForm("key")
	dstPath := path.Join(DefaultDir, container, key)
	dir := path.Dir(dstPath)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	ctx.LogBuffer.WriteLog("file_name[%s], dst_path[%s], dir[%s]", file.Filename, dstPath, dir)
	err = c.SaveUploadedFile(file, dstPath)
	if err != nil {
		ctx.LogBuffer.WriteLog("save_file[failed] error_msg[%s]", err.Error())
		return
	}

	res.FileName = file.Filename
}

type DownloadReq struct {
	Container string `json:"container" binding:"required"`
	Key       string `json:"key" binding:"required"`
}

type DownloadRes struct {
	middleware.CommonResponse
}

func Download(c *gin.Context) {
	ctx := context.GetContext()
	var res DownloadRes
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	var req DownloadReq
	if err = middleware.GetJsonParam(c, &req); err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}

	dst := path.Join(DefaultDir, req.Container, req.Key)
	ctx.LogBuffer.WriteLog("dst[%s]", dst)
	c.File(dst)
}
