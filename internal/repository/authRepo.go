package repository

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) auth.UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) GetByEmail(ctx context.Context, email string) (*auth.User, error) {
	var user auth.User
	err := u.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return &user, err
}

func (u *userRepo) GetById(ctx context.Context, id uint) (*auth.User, error) {
	var user auth.User
	err := u.db.WithContext(ctx).Where("id=?", id).First(&user).Error
	return &user, err
}

func (u *userRepo) Create(ctx context.Context, user *auth.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userRepo) Update(ctx context.Context, user *auth.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}
