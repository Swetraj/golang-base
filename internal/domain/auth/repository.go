package auth

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetById(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, user *User) error
}

type VerificationTokenRepository interface {
	Create(ctx context.Context, token *VerificationToken) error
	GetByToken(ctx context.Context, token string) (*VerificationToken, error)
	Update(ctx context.Context, token *VerificationToken) error
}
