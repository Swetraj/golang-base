package handler

import (
	"context"
	"github.com/Swetraj/golang-base/api/v1/dto"
	"github.com/Swetraj/golang-base/internal/pkg/validations"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (handler *RoutesHandler) RegisterUser(c *gin.Context) {
	var userRequest dto.RegisterRequest

	if !validations.BindAndValidate(c, &userRequest) {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := handler.userService.Register(ctx, userRequest.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to register user"})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"message": "Successfully Register User",
		},
	)
}
