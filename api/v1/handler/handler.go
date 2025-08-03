package handler

import (
	"github.com/Swetraj/golang-base/internal/domain/auth"
	"github.com/Swetraj/golang-base/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseHandler struct {
	validator *validator.Validate
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		validator: validator.New(),
	}
}

type AuthHandler struct {
	*BaseHandler
	userService  auth.UserService
	tokenService auth.VerificationService
}

func NewHandler(services *services.Services) *AuthHandler {
	return &AuthHandler{
		BaseHandler:  NewBaseHandler(),
		userService:  services.Auth,
		tokenService: services.Token,
	}
}

func (handler *AuthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	hg := rg.Group("/auth")
	{
		hg.POST("/login", handler.Login)
		hg.POST("/register", handler.RegisterUser)
		hg.POST("/reset", handler.ResetPwd)
	}
}
