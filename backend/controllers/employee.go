package controllers

import (
	"context"
	"employeeperformance/config"
	"employeeperformance/models"
	"employeeperformance/responses"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection for employees in MongoDB
var employeeCollection *mongo.Collection = config.GetCollection(config.DB, "employees")

// Validator instance for validating employee fields
var validate = validator.New()

// CreateEmployee handles the creation of a new employee in the database
func CreateEmployee(c *fiber.Ctx) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var employee models.Employee

	// Parse the request body into the employee struct
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.EmployeeResponse{
				Status: http.StatusBadRequest,
				Message: "Error parsing request body",
				Data: &fiber.Map{"data": err.Error()},
			})
	}

	// Validate required employee fields
	if validationErr := validate.Struct(&employee); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.EmployeeResponse{
				Status: http.StatusBadRequest,
				Message: "Error in employee data",
				Data: &fiber.Map{"data": validationErr.Error()},
			})
	}

	// Seed random number generator and create new Employee
	rand.Seed(time.Now().UnixNano())
	newEmployee := models.Employee{
		Id:         rand.Int63(),
		Name:       employee.Name,
		Performance: employee.Performance,
		Date:       time.Now(),
	}

	// Insert new employee into the database
	result, err := employeeCollection.InsertOne(ctx, newEmployee)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.EmployeeResponse{
				Status: http.StatusInternalServerError,
				Message: "Error inserting employee into database",
				Data: &fiber.Map{"data": err.Error()},
			})
	}

	// Return success response with created employee data
	return c.Status(http.StatusCreated).JSON(
		responses.EmployeeResponse{
			Status: http.StatusCreated,
			Message: "Employee successfully created",
			Data: &fiber.Map{"data": result},
		})
}

// GetAllEmployees handles fetching all employees from the database
func GetAllEmployees(c *fiber.Ctx) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var employees []models.Employee

	// Find all employees in the database
	results, err := employeeCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.EmployeeResponse{
				Status: http.StatusInternalServerError,
				Message: "Error fetching employees from database",
				Data: &fiber.Map{"data": err.Error()},
			})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleEmployee models.Employee
		if err = results.Decode(&singleEmployee); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.EmployeeResponse{
					Status: http.StatusInternalServerError,
					Message: "Error decoding employee data",
					Data: &fiber.Map{"data": err.Error()},
				})
		}
		employees = append(employees, singleEmployee)
	}

	// Return success response with employees data
	return c.Status(http.StatusOK).JSON(
		responses.EmployeeResponse{
			Status: http.StatusOK,
			Message: "Successfully fetched employees",
			Data: &fiber.Map{"data": employees},
		})
}
