package repository

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"gorm.io/gorm"
)

type Repositories struct {
	Auth  auth.UserRepository
	Token auth.VerificationTokenRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Auth:  NewUserRepository(db),
		Token: NewTokenRepository(db),
	}
}
