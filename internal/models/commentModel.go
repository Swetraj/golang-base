package models

import (
	"github.com/Swetraj/golang-base/internal/models/user"
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"foreignkey:PostID" json:"postID" binding:"required,gt=0"`
	UserID    uint   `gorm:"foreignkey:UserID"`
	Body      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User
}
