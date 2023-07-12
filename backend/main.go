package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"employeeperformance/config"
	"employeeperformance/routes"
)

func main() {
	// Connect to database
	config.Connect()
	
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Apply the middleware to the router (use cors.Default() to allow all origins)
	router.Use(cors.Default())

	// Setup route group for the API
	routes.Routes(router)

	// Start and run the server
	router.Run(":5001")
}
