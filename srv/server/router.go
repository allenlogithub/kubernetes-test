package server

import (
	"github.com/gin-gonic/gin"

	"srv/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		sayGroup := v1.Group("say")
		{
			say := new(controllers.Controllers)
			sayGroup.GET("/hi", say.SayHi)
		}
		postGroup := v1.Group("post")
		{
			post := new(controllers.Controllers)
			postGroup.POST("/content", post.PostContent)
		}
		getGroup := v1.Group("get")
		{
			get := new(controllers.Controllers)
			getGroup.GET("/allcontent", get.GetAllContent)
		}
	}

	return router
}
