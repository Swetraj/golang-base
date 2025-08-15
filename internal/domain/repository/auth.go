package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetById(ctx context.Context, id uint) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}

type VerificationTokenRepository interface {
	Create(ctx context.Context, token *model.VerificationToken) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, token *model.VerificationToken) error
	GetByToken(ctx context.Context, token string) (*model.VerificationToken, error)
	Update(ctx context.Context, token *model.VerificationToken) error
}
