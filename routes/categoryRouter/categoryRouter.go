package categoryRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/controllers/category"
  "go-server/middleware"
)

func SetupRouter (router *gin.Engine) {
	categoryRouter := router.Group("/category")

	categoryRouter.GET("/all", category.GetAll)

  categoryRouter.PATCH("/:old/:new", middleware.RequireAdmin(), category.ChangeCategoryName)
}
