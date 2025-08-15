package services

import (
	"github.com/Swetraj/golang-base/internal/domain/service"
	"github.com/Swetraj/golang-base/internal/repository"
	"gorm.io/gorm"
)

type Services struct {
	Auth  service.UserService
	Token service.VerificationService
}

func NewServices(repos *repository.Repositories, db *gorm.DB) *Services {
	return &Services{
		Auth:  NewUserService(db, repos.Auth, repos.Token),
		Token: NewTokenService(repos.Token),
	}

}
