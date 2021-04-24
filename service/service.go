package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/errcode"
	"github.com/glstr/futty_golang/logger"
)

type Service interface {
	LoadService(g *gin.RouterGroup) error
}

var DefaultServices map[string]Service = map[string]Service{
	"/snow": &homeService,
	"/cmd":  &cmdService,
	"/file": &fileOp,
}

type CommonRes struct {
	errcode.ErrorInfo
}

func guardCallback(c *gin.Context, ctx *context.Context) {
	if err, ok := recover().(error); ok {
		ctx.LogBuffer.WriteLog("[error_msg:%s]", err.Error())
		res := &CommonRes{
			errcode.InternalError,
		}
		c.JSON(http.StatusBadRequest, res)
	}
	logger.Notice(ctx.LogBuffer.String())
}
