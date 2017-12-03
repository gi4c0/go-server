package articleRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/middleware"
)

func setupRouter (router *gin.Engine) *gin.RouterGroup {
	articleRouter := router.Group("/article")

	articleRouter.POST("/", middleware.RequireAuth())

	return articleRouter
}