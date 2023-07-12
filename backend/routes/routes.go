package routes

import (
	"employeeperformance/controllers"

	"github.com/gofiber/fiber/v2"
)

// EmployeeRoutes function sets up the routes for employee operations
func EmployeeRoutes(app *fiber.App) {
	// POST request to /employees will create a new employee
	app.Post("/employees", controllers.CreateEmployee)

	// GET request to /employees will fetch all employees
	app.Get("/employees", controllers.GetAllEmployees)
}
