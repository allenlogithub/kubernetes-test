package server

import (
	"github.com/gin-gonic/gin"

	controllers "srv/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		hiGroup := v1.Group("say")
		{
			hi := new(controllers.Controllers)
			hiGroup.GET("/hi", hi.SayHi)
			hiGroup.GET("/hihi", hi.SayHiHi)
		}
	}

	return router
}
