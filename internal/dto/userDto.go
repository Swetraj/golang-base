package dto

import (
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"time"
)

type UserResponseDTO struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SingleUserToUserResponseDTO(user userModel.User) UserResponseDTO {
	dto := UserResponseDTO{
		ID:        user.ID,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return dto
}

func UsersToUserResponseDTOs(users []userModel.User) ([]UserResponseDTO, error) {
	dtos := make([]UserResponseDTO, len(users))
	for i, user := range users {
		dtos[i] = UserResponseDTO{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return dtos, nil
}
