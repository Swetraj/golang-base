package auth

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/auth"
)

func (u *userRepo) Update(ctx context.Context, user *auth.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}

func (t *tokenRepo) Update(ctx context.Context, token *auth.VerificationToken) error {
	return t.db.WithContext(ctx).Save(token).Error
}
