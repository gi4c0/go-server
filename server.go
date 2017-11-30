package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"test/routes"
)


func main() {
	app := iris.New()

	app.Use(logger.New())

	app.Controller("/user", new(routes.UserController))

	app.Run(iris.Addr(":8000"))
}


