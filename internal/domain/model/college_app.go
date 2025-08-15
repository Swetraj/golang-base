package model

import (
	"gorm.io/gorm"
	"time"
)

type CollegeApplication struct {
	gorm.Model
	StudentID      uint
	Status         string
	DecisionDate   time.Time
	SubmissionDate time.Time
	Student        Student `gorm:"foreignKey:StudentID"`
}

type SATDetails struct {
	gorm.Model
	StudentID uint
	Math      uint
	English   uint
	Essay     uint
	Student   Student `gorm:"foreignKey:StudentID"`
}

type EnglishProficiency struct {
	gorm.Model
	Name      string
	Score     uint
	StudentID uint
	Student   Student `gorm:"foreignKey:StudentID"`
}
