package model

import "gorm.io/gorm"

type ResourcePage struct {
	gorm.Model
	Title         string
	Subtitle      string
	IntroTitle    string
	IntroSubtitle string
	IsActive      bool
	IsDraft       bool
}
