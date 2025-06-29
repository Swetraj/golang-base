package tests

import (
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/models"
	"github.com/Swetraj/golang-base/internal/models/user"
	"github.com/joho/godotenv"
	"log"
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
	err = initializers.DB.Migrator().DropTable(user.User{}, models.Category{}, models.Post{}, models.Comment{})
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	// Migrate again
	err = initializers.DB.AutoMigrate(user.User{}, models.Category{}, models.Post{}, models.Comment{})

	if err != nil {
		log.Fatal("Migration failed")
	}
}
