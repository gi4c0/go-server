package userRouter

import (
	"github.com/gin-gonic/gin"
	"test/controllers/auth"
)

func SetupRouter(router *gin.Engine) {
	userRouter := router.Group("/user")

	userRouter.POST("/register", auth.Register)
	userRouter.POST("/login", auth.Login)
	userRouter.GET("/auth", auth.Auth)
}
