package main

import (
	"github.com/Library-API-CLI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// set up the router
	router := gin.Default()

	// set up routes
	routes.SetUpRoutes(router)

	// start the server
	router.Run("localhost:8080")
}