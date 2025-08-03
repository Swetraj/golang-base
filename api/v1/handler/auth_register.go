package handler

import (
	"context"
	"errors"
	"github.com/Swetraj/golang-base/api/v1/dto"
	"github.com/Swetraj/golang-base/internal/pkg/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func (handler *RoutesHandler) RegisterUser(c *gin.Context) {
	var userRequest dto.RegisterRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			c.JSON(
				http.StatusUnprocessableEntity, gin.H{
					"validations": validations.FormatValidationErrors(errs),
				},
			)
			return
		}

		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if err := handler.validator.Struct(&userRequest); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			c.JSON(
				http.StatusUnprocessableEntity, gin.H{
					"validations": validations.FormatValidationErrors(errs),
				},
			)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
