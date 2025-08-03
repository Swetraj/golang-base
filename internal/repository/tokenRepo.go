package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"gorm.io/gorm"
	"time"
)

type tokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) auth.VerificationTokenRepository {
	return &tokenRepo{db: db}
}

func (t *tokenRepo) Create(ctx context.Context, token *auth.VerificationToken) error {
	return t.db.WithContext(ctx).Create(token).Error
}

func (t *tokenRepo) GetByToken(ctx context.Context, tokenString string) (*auth.VerificationToken, error) {
	var token auth.VerificationToken
	err := t.db.WithContext(ctx).Where("token=?", tokenString).Where("expiry >", time.Now()).First(&token).Error
	return &token, err
}

func (t *tokenRepo) Update(ctx context.Context, token *auth.VerificationToken) error {
	return t.db.WithContext(ctx).Save(token).Error
}
