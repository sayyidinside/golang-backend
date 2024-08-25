package database

import (
	"fmt"
	"golang/backend/models"
	"golang/backend/utils"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {
	// Load database credentials from environment variables
	dbHost := utils.GetENVWithDefault("DB_HOST", "localhost")
	dbPort := utils.GetENVWithDefault("DB_PORT", "3306")
	dbUser := utils.GetENVWithDefault("DB_USERNAME", "root")
	dbPassword := utils.GetENVWithDefault("DB_PASSWORD", "")
	dbName := utils.GetENVWithDefault("DB_NAME", "backend")

	// Construct database URL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		os.Exit(2)
	}

	DBConn = db
	fmt.Println("Database succefully connected")

	// Migrate model
	migrate()
}

func migrate() {
	DBConn.AutoMigrate(&models.ProductCategory{})
	DBConn.AutoMigrate(&models.Product{})

	fmt.Println("Migration ruccessfully executed")
}
