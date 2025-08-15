package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *model.Profile) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, profile *model.Profile) error
	GetProfileByID(ctx context.Context, profileId uint) (*model.Profile, error)
	GetByEmail(ctx context.Context, email string) (*model.Profile, error)
	Update(ctx context.Context, user *model.Profile) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, profile *model.Profile) error
}
