package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type ServiceRepository interface {
	Create(ctx context.Context, serviceCat *model.ServiceCategory) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, serviceCat *model.ServiceCategory) error
	GetById(ctx context.Context, id uint) (*model.ServiceCategory, error)
	Update(ctx context.Context, serviceCat *model.ServiceCategory) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, serviceCat *model.ServiceCategory) error
}
