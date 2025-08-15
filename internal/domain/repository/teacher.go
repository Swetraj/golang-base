package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type TeacherRepository interface {
	Create(ctx context.Context, teacher *model.Teacher) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, teacher *model.Teacher) error
	GetByProfileID(ctx context.Context, profileId uint) (*model.Teacher, error)
	GetByEmail(ctx context.Context, email string) (*model.Teacher, error)
	Update(ctx context.Context, teacher *model.Teacher) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, teacher *model.Teacher) error
}
