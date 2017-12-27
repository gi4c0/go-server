package articleRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/middleware"
	"go-server/controllers/article"
	"go-server/controllers/comment"
)

func SetupRouter (router *gin.Engine) {
	articleRouter := router.Group("/article")
	commentRouter := articleRouter.Group("/comment")

	// Article Router
	articleRouter.GET("/id/:id", article.GetSingleArticle)
	articleRouter.GET("/list/:page/:count", article.GetArticles)
	articleRouter.GET("/my", middleware.RequireAuth(), article.GetUserArticlesPreview)
  articleRouter.GET("/unapproved", middleware.RequireAdmin(), article.GetUnapproved)

	articleRouter.POST("/", middleware.RequireAuth(), article.CreateArticle)

	articleRouter.PATCH("/update/:id", middleware.RequireAuth(), article.UpdateArticle)
	articleRouter.PATCH("/approve/:id", middleware.RequireAdmin(), article.ApproveArticle)

	articleRouter.DELETE("/image/:id", middleware.RequireAuth(), article.DeleteImage)

	// Comment Router
	commentRouter.GET("/:articleId", comment.GetComments)

	commentRouter.POST("/:articleId", middleware.RequireAuth(), comment.AddComment)

	commentRouter.PATCH("/:commentId", middleware.RequireAuth(), comment.UpdateComment)

	commentRouter.DELETE("/:commentId", middleware.RequireAuth(), comment.DeleteComment)
}
