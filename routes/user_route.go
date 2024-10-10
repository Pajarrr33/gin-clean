package routes

import (
	"gin-db/controller"
	"gin-db/middleware"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.Engine, uc controller.UserController) {
	userRoutes := router.Group("/api/v1/users").Use(middleware.BasicAuth())
	{
		userRoutes.POST("/",uc.CreateUser)
	}
}