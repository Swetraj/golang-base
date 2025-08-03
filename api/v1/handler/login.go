package handler

import (
	"context"
	"errors"
	"github.com/Swetraj/golang-base/api/v1/dto"
	"github.com/Swetraj/golang-base/internal/pkg/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func (handler *AuthHandler) Login(c *gin.Context) {
	var userRequest dto.LoginRequest

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

	user, err := handler.userService.Login(ctx, userRequest.Email, userRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tknString, err := handler.createJWT(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to authorize user"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tknString, 3600*24*30, "", "", false, true)

	c.JSON(
		http.StatusOK, gin.H{},
	)
}

func (handler *AuthHandler) createJWT(id uint, email string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": id,
			"email":  email,
			"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
