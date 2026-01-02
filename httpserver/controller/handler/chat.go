package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/service"
)

type ChatRequest struct {
	Service string `json:"service"`
	Message string `json:"message"`
}

type ChatResponse struct {
	middleware.CommonResponse
	Response string `json:"response"`
}

func ChatHandler(c *gin.Context) {
	ctx := context.GetContext()
	var res ChatResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	var req ChatRequest
	err = middleware.GetJsonParam(c, &req)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("message[%s]", req.Message)

	//call service
	ser, err := service.GetChatService(req.Service)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_service[failed] error_msg[%s]", err.Error())
		return
	}

	response, err := ser.Chat(ctx.LogBuffer, req.Message)
	if err != nil {
		ctx.LogBuffer.WriteLog("chat[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("response[%s]", response)

	res.Response = response
}
