package categoryRouter

import (
	"github.com/gin-gonic/gin"
	"go-server/controllers/category"
)

func SetupRouter (router *gin.Engine) {
	categoryRouter := router.Group("/category")

	categoryRouter.GET("/all", category.GetAll)
}
