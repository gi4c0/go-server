package articleRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/middleware"
	"go-server/controllers/article"
	"go-server/controllers/comment"
)

func SetupRouter (router *gin.Engine) {
	articleRouter := router.Group("/article")
	commentRouter := articleRouter.Group("/comments")

	articleRouter.GET("/:page/:count", article.GetArticles)

	articleRouter.POST("/", middleware.RequireAuth(), article.CreateArticle)

	articleRouter.PATCH("/:id", middleware.RequireAuth(), article.UpdateArticle)
	articleRouter.DELETE("/image/:id", middleware.RequireAuth(), article.DeleteImage)

	// Comment Router
	commentRouter.POST("/:articleId", middleware.RequireAuth(), comment.AddComment)
}