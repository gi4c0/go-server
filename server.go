package main

import (
	"test/routes"
)


func main() {
	router := routes.SetupRouter()

	router.Run(":8000")
}


