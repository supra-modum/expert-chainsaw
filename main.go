package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"expert-chainsaw/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	checkConnection()

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/users", func(c *gin.Context) {
			user.GetUsers(c, db)
		})
		api.POST("/users", func(c *gin.Context) {
			user.CreateUser(c, db)
		})
		api.DELETE("/users/:id", func(c *gin.Context) {
			user.DeleteUser(c, db)
		})
		api.PUT("/users/:id", func(c *gin.Context) {
			user.UpdateUser(c, db)
		})
	}

	router.Run(":8080")
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return db, nil
}

func checkConnection() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Create the connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database")
}
