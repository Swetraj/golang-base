package model

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	FullName    string
	Address     string
	Phone       string
	Gender      string
	DateOfBirth *time.Time
	ProfileType string `gorm:"not null"`
	UserID      uint   `gorm:"uniqueIndex"`
	User        User   `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
}
