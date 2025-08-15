package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(ctx context.Context, session *model.Session) error
	CreateWithTx(ctx context.Context, tx *gorm.DB, session *model.Session) error
	GetById(ctx context.Context, id uint) (*model.Session, error)
	Update(ctx context.Context, session *model.Session) error
	UpdateWithTx(ctx context.Context, tx *gorm.DB, session *model.Session) error
}
