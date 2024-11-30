package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB global variable to hold the DB connection
var DB *gorm.DB

// ConnectionDatabase initializes the connection to the MySQL database and runs migrations
func ConnectionDatabase() {
	// Load the .env file containing environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read the database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Ensure that necessary environment variables are set
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatalf("Missing required database environment variables")
	}

	// Create a Data Source Name (DSN) for MySQL
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Run auto migrations for the defined models
	// Ensure that your models (BahasaPasien, Kelurahan) are correctly defined in separate files
	if err := db.AutoMigrate(&BahasaPasien{}, &Kelurahan{}); err != nil {
		log.Fatalf("Error during auto-migration: %v", err)
	}

	// Assign the database connection to the global DB variable
	DB = db

	// Optionally, you can defer closing the database connection if required
	// dbSQL, _ := db.DB()
	// defer dbSQL.Close()
}
