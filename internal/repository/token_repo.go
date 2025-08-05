package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"github.com/Swetraj/golang-base/internal/domain/repository"
	"gorm.io/gorm"
)

type tokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) repository.VerificationTokenRepository {
	return &tokenRepo{db: db}
}

func (t *tokenRepo) Create(ctx context.Context, token *model.VerificationToken) error {
	return t.db.WithContext(ctx).Create(token).Error
}

func (t *tokenRepo) GetByToken(ctx context.Context, tokenString string) (*model.VerificationToken, error) {
	var token model.VerificationToken
	err := t.db.WithContext(ctx).Where("token=?", tokenString).First(&token).Error
	return &token, err
}

func (t *tokenRepo) Update(ctx context.Context, token *model.VerificationToken) error {
	return t.db.WithContext(ctx).Save(token).Error
}
