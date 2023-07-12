package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"employeeperformance/config"
	"employeeperformance/routes"
)

// Server port
const ServerPort = ":5001"

func main() {
	// Initialize Fiber application
	app := fiber.New()

	// Use CORS middleware to allow Cross-Origin Resource Sharing
	app.Use(cors.New())

	// Connect to MongoDB database
	config.ConnectDB()

	// Setup route group for the API
	routes.EmployeeRoutes(app)

	// Listen and serve the application on the defined port
	app.Listen(":5001")
}
