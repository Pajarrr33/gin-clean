package routes

import (
	"gin-db/controller"

	"github.com/gin-gonic/gin"
)

func Credentials(router *gin.Engine, cc controller.CredentialController) {
	credentialRoutes := router.Group("/api/v1")
	{
		credentialRoutes.POST("/register",cc.Register)
	}
}