package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"employeeperformance/config"
	"employeeperformance/routes"
)

func main() {
	// Initialize app
	app := fiber.New()

	// Add middleware
	app.Use(cors.New())

	// Connect to database
	config.ConnectDB()

	// Setup route group for the API
	routes.EmployeeRoutes(app)

	// Start and run the server
	app.Listen(":5001")
}
