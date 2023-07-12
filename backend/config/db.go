package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"

	"employeeperformance/controllers"
)

func Connect() *pg.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opts := &pg.Options{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Database: os.Getenv("DATABASE_NAME"),
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")

	controllers.CreateEmployeeTable(db)
	controllers.InitiateDB(db)
	
	return db
}
