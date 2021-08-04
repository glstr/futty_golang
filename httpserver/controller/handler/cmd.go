package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/model/service"
)

type CmdRequest struct {
	Method string `json:"method"`
	Args   string `json:"args"`
}

type CmdResponse struct {
	middleware.CommonResponse
	Result string
}

func Cmd(c *gin.Context) {
	ctx := context.GetContext()
	var res CmdResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(http.StatusOK, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	//get param
	var req CmdRequest
	err = middleware.GetJsonParam(c, &req)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("method[%s] ", req.Method)

	ser := service.GetCmdService()
	var result string
	result, err = ser.Exec(req.Method, req.Args)
	if err != nil {
		ctx.LogBuffer.WriteLog("exec[failed] error_msg[%s]", err.Error())
	}
	res.Result = result
}
