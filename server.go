package main

import (
	"go-server/routes"
)


func main() {
	router := routes.SetupRouter()

	router.Run(":8000")
}


