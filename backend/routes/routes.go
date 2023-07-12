package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"employeeperformance/controllers"
)

//Routes for API
func Routes(router *gin.Engine) {
	// Get all employees
	router.GET("/employees", controllers.GetAllEmployees)

	// Create an employee
	router.POST("/employees", controllers.CreateEmployee)

	// Get a single employee
	// router.GET("/employees/:employeeID", controllers.GetSingleEmployee)

	// Edit an employee
	// router.PUT("/employees/:employeeID", controllers.EditEmployee)

	// Delete an employee
	// router.DELETE("/employees/:employeeID", controllers.DeleteEmployee)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Employee Performance API",
	})
	return
}
