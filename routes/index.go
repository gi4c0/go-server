package routes

import (
	"github.com/gin-gonic/gin"
	"test/routes/userRouter"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRouter.SetupRouter(router)

	return router
}
