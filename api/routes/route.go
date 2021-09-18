package routes

import (
	"gin-derived/api/controller"
	"gin-derived/api/middleware"
	ws "gin-derived/pkg/websocket"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()

	//ws地址
	router.GET("/ws", ws.WsHandler)
	//使用跨域
	router.Use(middleware.Cors())
	//IM路由
	imGroup := router.Group("/im")
	{
		imGroup.POST("/login", controller.ImLogin)
	}
	//使用jwt
	imGroup.Use(middleware.JWT())
	{
		imGroup.POST("/get_user_contacts", controller.GetUserContacts)
		imGroup.POST("/get_messages", controller.GetMessages)
		imGroup.POST("/upload_avatar", controller.UploadAvatar)
		imGroup.POST("/update_userinfo", controller.UpdateUserinfo)
		imGroup.POST("/upload_chat_image", controller.UploadChatImage)
		imGroup.POST("/create_group", controller.CreateGroup)
		imGroup.POST("/get_group_contacts", controller.GetGroupContacts)
		imGroup.POST("/add_group_members", controller.AddGroupMembers)
		imGroup.POST("/get_group_members", controller.GetGroupMembers)
		imGroup.POST("/upload_group", controller.UploadGroup)
	}

	//普通接口
	group := router.Group("/")
	group.GET("/index", controller.Index)
	group1 := group.Group("/jwt", middleware.JWT())
	{
		group1.GET("/index", controller.Index)
	}
	return router
}
