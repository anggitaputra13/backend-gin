package tests

import (
	"log"

	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"github.com/joho/godotenv"
)

// DatabaseRefresh runs fresh migration
func DatabaseRefresh() {
	// Load env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	initializers.ConnectDB()

	// Drop all the tables
	err = initializers.DB.Migrator().DropTable(models.User{}, models.Order{}, models.Book{}, models.Employee{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	// Migrate again
	err = initializers.DB.AutoMigrate(models.User{}, models.Order{}, models.Book{}, models.Employee{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
