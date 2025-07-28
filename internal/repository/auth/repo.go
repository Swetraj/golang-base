package auth

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

type tokenRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) auth.UserRepository {
	return &userRepo{db: db}
}

func NewTokenRepository(db *gorm.DB) auth.VerificationTokenRepository {
	return &tokenRepo{db: db}
}
