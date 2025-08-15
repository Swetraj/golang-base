package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	School         string
	Level          string
	EnrollmentDate *time.Time
	ProfileID      uint
	StudentQueryID uint
	Profile        Profile
	StudentQuery   StudentQuery
	Parent         []Parent `gorm:"many2many:parent_student_detail"`
}

type StudentQuery struct {
	gorm.Model
	ServiceRequired string
	Subjects        pq.StringArray `gorm:"type:text[]"`
	PreferredMode   string
	Timing          string
	Found           string
	Query           string
}
