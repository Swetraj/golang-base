package model

import "gorm.io/gorm"

type ServiceCategory struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Levels      []Level
	Subjects    []Subject
}

type Level struct {
	gorm.Model
	ServiceCategoryID uint   `gorm:"not null"`
	Name              string `gorm:"not null"`
	DisplayName       string
	Description       string
	OrderIndex        uint
	ServiceCategory   ServiceCategory `gorm:"foreignKey:ServiceCategoryID"`
}

type Subject struct {
	gorm.Model
	LevelID           uint   `gorm:"not null"`
	ServiceCategoryID uint   `gorm:"not null"`
	Name              string `gorm:"not null"`
	Description       string
	ServiceCategory   ServiceCategory `gorm:"foreignKey:ServiceCategoryID"`
	Level             Level           `gorm:"foreignKey:LevelID"`
}
