package main

import (
	"github.com/Swetraj/golang-base/config"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"log"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	err := initializers.DB.AutoMigrate(
		auth.User{},
		auth.VerificationToken{},
	)

	if err != nil {
		log.Fatal("Migration failed")
	}
}
