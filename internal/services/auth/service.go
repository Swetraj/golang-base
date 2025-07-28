package auth

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
)

type tokenService struct {
	repo auth.VerificationTokenRepository
}

type userService struct {
	repo         auth.UserRepository
	tokenService tokenService
}

func NewUserService(repo auth.UserRepository, tokenService tokenService) auth.UserService {
	return &userService{repo, tokenService}
}

func NewTokenService(repo auth.VerificationTokenRepository) auth.VerificationService {
	return &tokenService{repo}
}
