package main

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
	repo "github.com/Swetraj/golang-base/internal/repository/auth"
	"gorm.io/gorm"
)

type Repositories struct {
	Auth  auth.UserRepository
	Token auth.VerificationTokenRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Auth: repo.NewUserRepository(db),
	}
}
