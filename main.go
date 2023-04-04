package main

import (
	"database/sql"
	"expert-chainsaw/models"
	secret "expert-chainsaw/tools"
	"expert-chainsaw/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"expert-chainsaw/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	secret.EnsureJwtSecret()

	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	checkConnection()

	router := gin.Default()

	api := router.Group("/api")

	{
		// User routes
		api.GET("/users", func(c *gin.Context) {
			handlers.GetUsers(c, db)
		})
		api.POST("/users", func(c *gin.Context) {
			handlers.CreateUser(c, db)
		})
		api.PUT("/users/:id", func(c *gin.Context) {
			handlers.UpdateUser(c, db)
		})
		api.POST("/login", func(c *gin.Context) {
			user.LoginUser(c, db)
		})
		api.DELETE("/users/:id", func(c *gin.Context) {
			handlers.DeleteUser(c, db)
		})

		// Fundraising routes
		api.GET("/fundraisings", func(c *gin.Context) {
			handlers.GetAllFundraisings(c, db)
		})
		api.POST("/fundraisings", func(c *gin.Context) {
			handlers.CreateFundraising(c, db)
		})
		api.PUT("/fundraisings/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
				return
			}
			handlers.UpdateFundraising(c, db, uint(id))
		})
		api.DELETE("/fundraisings/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
				return
			}
			handlers.DeleteFundraising(c, db, uint(id))
		})

		// User donation routes
		api.GET("/users/:user_id/donations", func(c *gin.Context) {
			handlers.GetUserDonations(c, db)
		})
		api.PUT("/donations/:id", func(c *gin.Context) {
			handlers.UpdateDonation(c, db)
		})
		api.POST("/donations", func(c *gin.Context) {
			handlers.AddDonation(c, db)
		})
	}

	router.Run(":8080")
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

	// Migrate the schema
	migrate(db)
	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Fundraising{}, &models.Donation{})

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
