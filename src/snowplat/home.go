package snowplat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SnowPlat struct {
	Router *gin.Engine
}

func NewSnowPlat(r *gin.Engine) *SnowPlat {
	return &SnowPlat{
		Router: r,
	}
}

func (s *SnowPlat) LoadRouter() {
	s.Router.LoadHTMLGlob("static/*")
	g := s.Router.Group("/snow")
	g.GET("/home.html", s.home)
}

func (s *SnowPlat) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "home",
	})
}
