package service

import (
	"net/http"

	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/errcode"
	"github.com/glstr/futty_golang/utils"

	"github.com/gin-gonic/gin"
)

type HomeService struct{}

var homeService HomeService

func (h *HomeService) LoadService(g *gin.RouterGroup) error {
	g.GET("/get_data", h.getTestData)
	return nil
}

type showMsgRes struct {
	CommonRes
	Msg []byte `json:"msg"`
}

type GetDataResponse struct {
	CommonRes
	Data utils.ShowData `json:"data"`
}

func (*HomeService) getTestData(c *gin.Context) {
	ctx := context.NewContext()
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
