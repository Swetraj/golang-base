package model

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	Name        string
	StartDate   *time.Time
	EndDate     *time.Time
	Description string
	Classes     []Class
}

type Class struct {
	gorm.Model
	SessionID    uint
	SubjectID    uint
	TeacherID    uint
	ClassCode    string `gorm:"unique;not null"`
	PlannedHours int
	ActualHours  int
	Session      Session
	Subject      Subject
	Teacher      Teacher `gorm:"foreignKey:TeacherID"`
}

type Enrollment struct {
	gorm.Model
	StudentID        uint
	ClassID          uint
	EnrollmentDate   *time.Time
	TuitionFeeStatus string
	Status           string
	Student          Student `gorm:"foreignKey:StudentID"`
	Class            Class   `gorm:"foreignKey:ClassID"`
}
