package auth

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/auth"
)

func (t *tokenService) UpdateToken(ctx context.Context, token *auth.VerificationToken) error {
	err := t.repo.Update(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
