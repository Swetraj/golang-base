package handler

import (
	"github.com/Swetraj/golang-base/api/middleware"
	"github.com/Swetraj/golang-base/internal/domain/service"
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

type RoutesHandler struct {
	*BaseHandler
	userService  service.UserService
	tokenService service.VerificationService
}

func NewHandler(services *services.Services) *RoutesHandler {
	return &RoutesHandler{
		BaseHandler:  NewBaseHandler(),
		userService:  services.Auth,
		tokenService: services.Token,
	}
}

func (handler *RoutesHandler) RegisterRoutes(rg *gin.RouterGroup) {
	hg := rg.Group("/auth")
	{
		hg.POST("/login", handler.Login)
		hg.POST("/register", handler.RegisterUser)
		hg.POST("/reset", handler.ResetPwd)
	}
	rg.Use(middleware.RequireAuth)
}
