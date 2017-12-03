package routes

import (
	"github.com/gin-gonic/gin"
	"go-server/routes/userRouter"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRouter.SetupRouter(router)

	return router
}
