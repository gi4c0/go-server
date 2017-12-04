package routes

import (
	"github.com/gin-gonic/gin"
	"go-server/routes/userRouter"
	"go-server/routes/articleRouter"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRouter.SetupRouter(router)
	articleRouter.SetupRouter(router)

	return router
}
