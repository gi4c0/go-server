package articleRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/middleware"
	"go-server/controllers/article"
)

func SetupRouter (router *gin.Engine) *gin.RouterGroup {
	articleRouter := router.Group("/article")

	articleRouter.GET("/:page/:count", article.GetArticles)

	articleRouter.POST("/", middleware.RequireAuth(), article.CreateArticle)

	articleRouter.PATCH("/:id", middleware.RequireAuth(), article.UpdateArticle)
	articleRouter.DELETE("/image/:id", middleware.RequireAuth(), article.DeleteImage)

	return articleRouter
}