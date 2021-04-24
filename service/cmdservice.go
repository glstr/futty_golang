package service

import (
	"net/http"

	"github.com/glstr/futty_golang/cmdhandler"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/errcode"

	"github.com/gin-gonic/gin"
)

type CmdService struct{}

var cmdService CmdService

func (c *CmdService) LoadService(g *gin.RouterGroup) error {
	g.GET("exec", c.Exec)
	return nil
}

type ExecRes struct {
	CommonRes
	Result string `json:"result"`
}

func (s *CmdService) Exec(c *gin.Context) {
	ctx := context.NewContext()
	logbuf := ctx.LogBuffer
	defer guardCallback(c, ctx)

	res := &ExecRes{
		CommonRes{
			errcode.OK,
		},
		"",
	}

	keyMethod := "method"
	keyParams := "params"

	method := c.Query(keyMethod)
	params := c.Query(keyParams)
	logbuf.WriteLog("[method:%s][params:%s]", method, params)

	var output string
	var err error
	if params == "" {
		output, err = cmdhandler.DefaultHandler.Execute(method)
	} else {
		output, err = cmdhandler.DefaultHandler.Execute(method, params)
	}
	if err != nil {
		res.CommonRes = CommonRes{errcode.InternalError}
		logbuf.WriteLog("[exec cmd fail][error_msg:%s]", err.Error())
		logbuf.WriteLog("[output:%s]", output)
		c.JSON(http.StatusOK, res)
		return
	}
	res.Result = output
	logbuf.WriteLog("[output:%s]", output)
	c.JSON(http.StatusOK, res)
}
