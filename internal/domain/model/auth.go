package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique;not null"`
	PasswordHash string `json:"-"`
	Provider     string `gorm:"not null;default:'email'" json:"provider"`
	IsVerified   bool   `gorm:"not null;default:false" json:"is_verified"`
	IsActive     bool   `gorm:"default:false"`
}

type VerificationToken struct {
	ID        uint   `gorm:"primarykey"`
	UserID    uint   `gorm:"not null"`
	Token     string `gorm:"uniqueIndex"`
	Type      string `gorm:"index"`
	ExpiresAt time.Time
	Used      bool
}
