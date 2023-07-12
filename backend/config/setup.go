package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Env file key for MongoDB connection string
const MongoDbURIKey = "MONGODB_URI"

// Database name to be used
const DatabaseName = "golangAPI"

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// ConnectDB establishes a connection with the MongoDB database
func ConnectDB() *mongo.Client {
	// Load environment variables
	LoadEnv()

	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv(MongoDbURIKey)))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Ensure the context cancellation is called to avoid memory leak
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the database to ensure connection establishment
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

// Global MongoDB client instance
var DB *mongo.Client = ConnectDB()

// GetCollection returns a collection from the database
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(DatabaseName).Collection(collectionName)
	return collection
}
