package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/service"
)

type GetMaterialListResponse struct {
	middleware.CommonResponse
	MaterialList []*service.MaterialInfo `json:"material_list"`
}

func GetMaterialList(c *gin.Context) {
	ctx := context.GetContext()
	var res GetMaterialListResponse
	res.CommonResponse.RequestId = ctx.Logid
	ctx.LogBuffer.WriteLog("method[GetMaterialList] ")
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		logger.Notice(ctx.LogBuffer.String())
		c.JSON(200, res)
		context.PutContext(ctx)
	}()

	//call service
	ser := service.GetMaterialService()
	var list []*service.MaterialInfo
	list, err = ser.GetMaterialList()
	if err != nil {
		ctx.LogBuffer.WriteLog("error_msg[%s]", err.Error())
		return
	}

	ctx.LogBuffer.WriteLog("res_len[%d]", len(list))
	res.MaterialList = list
}
