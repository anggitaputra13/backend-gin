package main

import (
	"log"

	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/config"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.Migrator().DropTable(models.User{}, models.Order{}, models.Book{}, models.Employee{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	err = initializers.DB.AutoMigrate(models.User{}, models.Order{}, models.Book{}, models.Employee{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
