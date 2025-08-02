package handler

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	userService auth.UserService
	validator   *validator.Validate
}

func NewHandler(userService auth.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		validator:   validator.New(),
	}
}
