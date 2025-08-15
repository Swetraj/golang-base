package model

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	gorm.Model
	HireDate       *time.Time
	EducationLevel string
	ProfileID      uint
	Profile        Profile
}
