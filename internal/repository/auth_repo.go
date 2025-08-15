package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
	"github.com/Swetraj/golang-base/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Where(
		&model.User{
			Email:    email,
			IsActive: true,
		},
	).First(&user).Error
	return &user, err
}

func (u *userRepo) GetById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Where("id=?", id).First(&user).Error
	return &user, err
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}
func (u *userRepo) CreateWithTx(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return tx.WithContext(ctx).Create(user).Error
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}
