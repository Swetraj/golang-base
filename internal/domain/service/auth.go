package service

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
)

type UserService interface {
	Register(ctx context.Context, email string) error
	Login(ctx context.Context, email string, password string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserById(ctx context.Context, id uint) (*model.User, error)
	ResetPassword(ctx context.Context, token string, pwd string) error
	SendEmail(email string, token string)
}

type VerificationService interface {
	UpdateToken(ctx context.Context, token *model.VerificationToken) error
}
