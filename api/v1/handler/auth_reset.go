package handler

import (
	"context"
	"github.com/Swetraj/golang-base/api/v1/dto"
	"github.com/Swetraj/golang-base/internal/pkg/validations"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (handler *RoutesHandler) ResetPwd(c *gin.Context) {

	var userRequest dto.ResetPasswordRequest
	query := c.Query("link")

	if !validations.BindAndValidate(c, &userRequest) {
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	err := handler.userService.ResetPassword(ctx, query, userRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"message": "Successfully reset password",
		},
	)
}
