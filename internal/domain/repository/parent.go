package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type ParentRepository interface {
	Create(ctx context.Context, parent *model.Parent) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, parent *model.Parent) error
	GetParentByStudentID(ctx context.Context, studentId uint) ([]*model.Parent, error)
}
