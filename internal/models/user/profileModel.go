package user

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	UserID         uint `gorm:"uniqueIndex"`
	ProfilePicture string
	FirstName      string
	MiddleName     string
	LastName       string
	Address        string
	Gender         string
	DateOfBirth    time.Time
	ProfileType    string `gorm:"not null"`
}
