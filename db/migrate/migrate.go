package main

import (
	"github.com/Swetraj/golang-base/config"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"log"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	err := initializers.DB.AutoMigrate(
		model.User{},
		model.VerificationToken{},
		model.Profile{},
		model.Parent{},
		model.StudentQuery{},
		model.Student{},
		model.Teacher{},
		model.ServiceCategory{},
		model.Level{},
		model.Subject{},
		model.CollegeApplication{},
		model.SATDetails{},
		model.EnglishProficiency{},
		model.HeroSectionCarousel{},
		model.HeroSection{},
		model.ImageCarouselImage{},
		model.ImageCarousel{},
		model.VideoSection{},
		model.ServiceSection{},
		model.ImpactSection{},
		model.HomePage{},
		model.ResourcePage{},
		model.ServicePage{},
		model.Session{},
		model.Class{},
		model.Enrollment{},
	)

	if err != nil {
		log.Fatal("Migration failed")
	}
}
