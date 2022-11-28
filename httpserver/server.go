package httpserver

import "github.com/gin-gonic/gin"

func StartHttpServer() error {
	e := gin.Default()
	err := LoadRouter(e)
	if err != nil {
		return err
	}
	e.Run(":8882")
	return nil
}
