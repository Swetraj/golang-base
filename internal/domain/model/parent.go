package model

import "gorm.io/gorm"

type Parent struct {
	gorm.Model
	ProfileID uint
	Profile   Profile `gorm:"constraint:OnDelete:CASCADE"`
}
