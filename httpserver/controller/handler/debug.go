package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/service"
)

type DebugRequest struct {
	Service string `json:"service"`
	Method  string `json:"method"`
}

type DebugResponse struct {
	middleware.CommonResponse
	Data interface{} `json:"data"`
}

func DebugHandler(c *gin.Context) {
	ctx := context.GetContext()
	var res DebugResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	ctx.LogBuffer.WriteLog("method[debug] ")
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	//get param & check param
	var req DebugRequest
	err = middleware.GetJsonParam(c, &req)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("service[%s] method[%s]", req.Service, req.Method)

	//call service or call services
	srv, err := service.GetDebugService(req.Service)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_service[failed], error_msg[%s]", err.Error())
		return
	}
	data, err := srv.Do(req.Method)
	if err != nil {
		ctx.LogBuffer.WriteLog("method_do[failed], error_msg[%s]", err.Error())
		return
	}

	//make resp
	res.Data = data
}
