package categoryRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/controllers/client/userCategory"
	"go-server/controllers/admin/adminCategory"
  "go-server/middleware"
)

func SetupRouter (router *gin.Engine) {
	categoryRouter := router.Group("/category")

	categoryRouter.GET("/all", userCategory.GetAll)

  categoryRouter.PATCH("/:old/:new", middleware.RequireAdmin(), adminCategory.ChangeCategoryName)
}
