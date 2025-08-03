package router

import (
	"github.com/gin-gonic/gin"
	"usercenter/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("register", controller.Register)
		user.POST("/login", controller.Login)
	}
	return router
}
