package service

import (
	"github.com/gin-gonic/gin"
)

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

func (s *SnowPlat) Load() {
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
