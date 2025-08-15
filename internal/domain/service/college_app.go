package service

import (
	"context"
	"github.com/Swetraj/golang-base/internal/domain/model"
)

type CollegeAppService interface {
	Initiate(ctx context.Context, collegeApp *model.CollegeApplication) error
	GetCollegeAppByStudentId(ctx context.Context, studentId uint) (*model.CollegeApplication, error)
	GetCollegeAppById(ctx context.Context, id uint) (*model.CollegeApplication, error)
	UpdateCollegeApp(ctx context.Context, collegeApp *model.CollegeApplication) error
}
