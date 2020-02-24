package snowplat

import (
	"errcode"
	"message"
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
)

type SnowPlat struct {
	Router *gin.Engine
}

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

func NewSnowPlat(r *gin.Engine) *SnowPlat {
	return &SnowPlat{
		Router: r,
	}
}

func (s *SnowPlat) LoadRouter() {
	g := s.Router.Group("/snow")
	g.GET("/show_msg", s.showmsg)
	g.GET("/get_data", s.getTestData)

	//file server
	s.Router.Static("/static", "./static/")
}

func (s *SnowPlat) home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"title": "post",
	})
}

type showMsgRes struct {
	CommonRes
	Msg []byte `json:"msg"`
}

func (s *SnowPlat) showmsg(c *gin.Context) {
	ctx := utils.NewContext()
	defer guardCallback(c, ctx)

	res := &showMsgRes{
		CommonRes{
			errcode.Ok,
		},
		[]byte(""),
	}

	msg, err := message.MakeDefaultMessage()
	if err != nil {
		ctx.LogBuffer.WriteLog("[error_msg:%s]", err.Error())
		res.CommonRes = CommonRes{errcode.InternalError}
		res.Msg = msg
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Msg = msg
	c.JSON(http.StatusOK, res)
	return
}

type GetDataResponse struct {
	CommonRes
	Data utils.ShowData `json:"data"`
}

func (*SnowPlat) getTestData(c *gin.Context) {
	ctx := utils.NewContext()
	defer guardCallback(c, ctx)

	showData := utils.GenerateData()
	res := &GetDataResponse{
		CommonRes{
			errcode.Ok,
		},
		showData,
	}
	c.JSON(http.StatusOK, res)
}
