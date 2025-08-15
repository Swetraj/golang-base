package model

import "gorm.io/gorm"

type ServicePage struct {
	gorm.Model
	Subtitle         string
	IntroductionBlog string `gorm:"type:jsonb"`
	TestPrepSubtitle string
	ExamPrepSubtitle string
	IsActive         bool
	IsDraft          bool
}
