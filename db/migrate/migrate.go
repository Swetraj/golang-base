package main

import (
	"github.com/Swetraj/golang-base/config"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/models"
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"log"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.Migrator().DropTable(
		userModel.User{},
		userModel.Permission{},
		userModel.PermissionCategory{},
		userModel.Role{},
		userModel.Profile{},
		models.Category{},
		models.Post{},
		models.Comment{},
	)
	if err != nil {
		log.Fatal("Table dropping failed")
	}

	err = initializers.DB.AutoMigrate(
		userModel.User{},
		userModel.Permission{},
		userModel.PermissionCategory{},
		userModel.Role{},
		userModel.Profile{},
		models.Category{},
		models.Post{},
		models.Comment{},
	)

	if err != nil {
		log.Fatal("Migration failed")
	}
}
