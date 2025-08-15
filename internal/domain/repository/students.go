package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(ctx context.Context, student *model.Student) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, student *model.Student) error
	GetByProfileID(ctx context.Context, profileId uint) (*model.Student, error)
	GetByEmail(ctx context.Context, email string) (*model.Student, error)
	Update(ctx context.Context, student *model.Student) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, student *model.Student) error
}

type StudentQuery interface {
	Create(ctx context.Context, query *model.StudentQuery) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, student *model.StudentQuery) error
	GetByProfileID(ctx context.Context, profileId uint) (*model.StudentQuery, error)
	Update(ctx context.Context, student *model.StudentQuery) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, student *model.StudentQuery) error
}
