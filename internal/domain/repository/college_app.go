package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type CollegeAppRepository interface {
	Create(ctx context.Context, collApp *model.CollegeApplication) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, collApp *model.CollegeApplication) error
	GetById(ctx context.Context, id uint) (*model.CollegeApplication, error)
	Update(ctx context.Context, collApp *model.CollegeApplication) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, collApp *model.CollegeApplication) error
}
