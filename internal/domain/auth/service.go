package auth

import "context"

type UserService interface {
	Register(ctx context.Context, email string) error
	Login(ctx context.Context, email string, password string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id uint) (*User, error)
	ResetPassword(ctx context.Context, token string, pwd string) error
}

type VerificationService interface {
	UpdateToken(ctx context.Context, token *VerificationToken) error
}
