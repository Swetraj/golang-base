package auth

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"time"
)

func (u *userRepo) GetByEmail(ctx context.Context, email string) (*auth.User, error) {
	var user auth.User
	err := u.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}

func (u *userRepo) GetById(ctx context.Context, id uint) (*auth.User, error) {
	var user auth.User
	err := u.db.WithContext(ctx).Where("id=?", id).First(&user).Error
	return &user, err
}

func (t *tokenRepo) GetByToken(ctx context.Context, tokenString string) (*auth.VerificationToken, error) {
	var token auth.VerificationToken
	err := t.db.WithContext(ctx).Where("token=?", tokenString).Where("expiry >", time.Now()).First(&token).Error
	return &token, err
}
