package routes

import (
	"github.com/gin-gonic/gin"
	"go-server/routes/userRouter"
	"go-server/routes/articleRouter"
	"go-server/routes/categoryRouter"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRouter.SetupRouter(router)
	articleRouter.SetupRouter(router)
	categoryRouter.SetupRouter(router)

	router.Static("/public", "./public")

	return router
}
