package services

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"github.com/Swetraj/golang-base/internal/repository"
)

type Services struct {
	Auth  auth.UserService
	Token auth.VerificationService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth:  NewUserService(repos.Auth, repos.Token),
		Token: NewTokenService(repos.Token),
	}

}
