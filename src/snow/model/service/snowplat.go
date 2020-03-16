package service

import (
	"errcode"
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
)

type CommonRes struct {
	errcode.ErrorInfo
}

func guardCallback(c *gin.Context, ctx *utils.Context) {
	if err, ok := recover().(error); ok {
		ctx.LogBuffer.WriteLog("[error_msg:%s]", err.Error())
		res := &CommonRes{
			errcode.InternalError,
		}
		c.JSON(http.StatusBadRequest, res)
	}
	ctx.Logger.Info(ctx.LogBuffer.String())
}

var DefaultServices map[string]Service = map[string]Service{
	"/snow": &homeService,
	"/cmd":  &cmdService,
}

type SnowPlat struct {
	Router   *gin.Engine
	services map[string]Service
}

func NewSnowPlat(r *gin.Engine) *SnowPlat {
	return &SnowPlat{
		Router:   r,
		services: DefaultServices,
	}
}

func (s *SnowPlat) LoadRouter() {
	for name, service := range s.services {
		g := s.Router.Group(name)
		err := service.LoadService(g)
		if err != nil {
			panic(err)
		}
	}
	//file server
	s.Router.Static("/static", "./static/")
}

type Service interface {
	LoadService(g *gin.RouterGroup) error
}
