package user

import (
	"gorm.io/gorm"
	"time"
)

type PermissionCategory struct {
	gorm.Model
	Order       int32  `gorm:"unique; not null"`
	Name        string `gorm:"unique; not null"`
	Description string `gorm:"not null"`
	Permission  []Permission
}

type Permission struct {
	gorm.Model
	SubOrder             int32   `gorm:"not null"`
	Description          string  `gorm:"not null"`
	PermissionCategoryID uint    `gorm:"not null"`
	Users                []*User `gorm:"many2many:user_permissions"`
	Roles                []*Role `gorm:"many2many:role_permissions"`
}

type Role struct {
	gorm.Model
	Name        string        `gorm:"not null"`
	Description string        `gorm:"not null"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
	Users       []*User       `gorm:"many2many:user_roles"`
}

type User struct {
	gorm.Model
	Email       string        `gorm:"unique;not null"`
	Username    string        `gorm:"unique;not null"`
	Password    string        `gorm:"unique;not null" json:"-"`
	IsActive    bool          `gorm:"default:false"`
	Permissions []*Permission `gorm:"many2many:user_permissions"`
	Roles       []*Role       `gorm:"many2many:user_roles"`
}

type PasswordReset struct {
	ID        uint   `gorm:"primarykey"`
	Email     string `gorm:"index"`
	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	Used      bool
}
