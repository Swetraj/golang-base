package repository

import (
	"github.com/Swetraj/golang-base/internal/domain/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	Auth  repository.UserRepository
	Token repository.VerificationTokenRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Auth:  NewUserRepository(db),
		Token: NewTokenRepository(db),
	}
}
