package articleRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/middleware"
	"go-server/controllers/article"
)

func SetupRouter (router *gin.Engine) *gin.RouterGroup {
	articleRouter := router.Group("/article")

	articleRouter.POST("/", middleware.RequireAuth(), article.CreateArticle)

	return articleRouter
}