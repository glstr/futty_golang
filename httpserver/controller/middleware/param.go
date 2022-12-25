package middleware

import "github.com/gin-gonic/gin"

func GetJsonParam(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindJSON(params); err != nil {
		return err
	}

	return nil
}

func GetParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBind(param); err != nil {
		return err
	}

	return nil
}
