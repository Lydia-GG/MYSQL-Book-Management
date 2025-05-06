package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	//  Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(" Error loading .env file")
	}

	//  Read credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	database, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to 'bookstore':", err)
	}
	db = database
}

// GetDB returns the initialized database instance

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is nil, ensure Connect() has been called.")
	}
	return db
}
