package articleRouter

import (
	"go-server/controllers/admin/adminArticle"
	"go-server/controllers/admin/adminComment"
	"go-server/controllers/client/userArticle"
	"go-server/controllers/client/userComment"
	"go-server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	articleRouter := router.Group("/article")
	commentRouter := articleRouter.Group("/comment")

	// Article Router
	articleRouter.GET("/id/:id", userArticle.GetSingleArticle)
	articleRouter.GET("/list/:page/:count", userArticle.GetArticles)
	articleRouter.GET("/my", middleware.RequireAuth(), userArticle.GetUserArticlesPreview)
	articleRouter.GET("/unapproved", middleware.RequireAdmin(), adminArticle.GetUnapproved)

	articleRouter.POST("/", middleware.RequireAuth(), userArticle.CreateArticle)

	articleRouter.PATCH("/update/:id", middleware.RequireAuth(), userArticle.UpdateArticle)
	articleRouter.PATCH("/approve/:id", middleware.RequireAdmin(), adminArticle.ApproveArticle)

	// articleRouter.DELETE("/image/:id", middleware.RequireAuth(), userArticle.DeleteImage)

	// Comment Router
	commentRouter.GET("/client/:articleId", userComment.GetComments)
	commentRouter.GET("/admin/:articleId", userComment.GetComments)

	commentRouter.POST("/admin/:articleId", middleware.RequireAuth(), adminComment.AddComment)
	commentRouter.POST("/client/:articleId", middleware.RequireAuth(), userComment.AddComment)

	commentRouter.PATCH("/:commentId", middleware.RequireAuth(), userComment.UpdateComment)

	commentRouter.DELETE("/:commentId", middleware.RequireAuth(), userComment.DeleteComment)
}
