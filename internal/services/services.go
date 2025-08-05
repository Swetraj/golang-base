package services

import (
	"github.com/Swetraj/golang-base/internal/domain/service"
	"github.com/Swetraj/golang-base/internal/repository"
)

type Services struct {
	Auth  service.UserService
	Token service.VerificationService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth:  NewUserService(repos.Auth, repos.Token),
		Token: NewTokenService(repos.Token),
	}

}
