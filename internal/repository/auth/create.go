package auth

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/auth"
)

func (u *userRepo) Create(ctx context.Context, user *auth.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (t *tokenRepo) Create(ctx context.Context, token *auth.VerificationToken) error {
	return t.db.WithContext(ctx).Create(token).Error
}
