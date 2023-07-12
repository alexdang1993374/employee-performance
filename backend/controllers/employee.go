package controllers

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

var dbConnect *pg.DB

// Initiate the database
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

type EmployeePerformance struct {
	ID        int64      `json:"id"`
	Name      string    `json:"name"`
	Performance  int    `json:"performance"`
	Date      time.Time `json:"Date"`
}

func CreateEmployeeTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&EmployeePerformance{}, opts)

	if createError != nil {
		log.Printf("Error while creating employee performance table, Reason: %v\n", createError)
		return createError
	}

	log.Printf("EmployeePerformance table created")

	return nil
}

func GetAllEmployees(c *gin.Context) {
	var employeePerformance []EmployeePerformance

	err := dbConnect.Model(&employeePerformance).Order("performance").Select()

	if err != nil {
		log.Printf("Error while getting all employees, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Employees",
		"data":    employeePerformance,
	})
	return
}

func CreateEmployee(c *gin.Context) {
	var employeePerformance EmployeePerformance
	c.BindJSON(&employeePerformance)

	name := employeePerformance.Name
	performance := employeePerformance.Performance

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63()

	insertError := dbConnect.Insert(&EmployeePerformance{
		ID:    id,
		Name:  name,
		Performance:   performance,
		Date: time.Now(),
	})

	if insertError != nil {
		log.Printf("Error while inserting new employee into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Employee created Successfully",
	})

	return
}
