package routes

import (
	"gin-derived/api/controller"
	"gin-derived/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	group := router.Group("/")
	group.GET("/index", controller.Index)
	group1 := group.Group("/jwt", middleware.JWT())
	{
		group1.GET("/index", controller.Index)
	}
	return router
}
