package service

import (
	"errcode"
	"message"
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
)

type HomeService struct{}

var homeService HomeService

func (h *HomeService) LoadService(g *gin.RouterGroup) error {
	g.GET("/show_msg", h.showmsg)
	g.GET("/get_data", h.getTestData)
	return nil
}

func (s *HomeService) home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"title": "post",
	})
}

type showMsgRes struct {
	CommonRes
	Msg []byte `json:"msg"`
}

func (s *HomeService) showmsg(c *gin.Context) {
	ctx := utils.NewContext()
	defer guardCallback(c, ctx)

	res := &showMsgRes{
		CommonRes{
			errcode.OK,
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

func (*HomeService) getTestData(c *gin.Context) {
	ctx := utils.NewContext()
	defer guardCallback(c, ctx)

	showData := utils.GenerateData()
	res := &GetDataResponse{
		CommonRes{
			errcode.OK,
		},
		showData,
	}
	c.JSON(http.StatusOK, res)
}
