package controllers

import (
	"errors"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/dto"
	"github.com/Swetraj/golang-base/internal/emails"
	"github.com/Swetraj/golang-base/internal/format-errors"
	"github.com/Swetraj/golang-base/internal/helpers"
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"github.com/Swetraj/golang-base/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Register Function is used to register email
func Register(c *gin.Context) {
	var userInput struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
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

	// Email unique validation
	if validations.IsUniqueValue("users", "email", userInput.Email) {
		c.JSON(
			http.StatusUnprocessableEntity, gin.H{
				"validations": map[string]interface{}{
					"Email": "The email is already exist!",
				},
			},
		)
		return
	}

	user := userModel.User{
		Email: userInput.Email,
	}

	// Create the user
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	token := helpers.GenerateResetToken()
	reset := userModel.PasswordReset{
		Token:     token,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	result = initializers.DB.Create(&reset)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	userResponse := dto.SingleUserToUserResponseDTO(user)

	c.JSON(
		http.StatusOK, gin.H{
			"user": userResponse,
		},
	)

	var email struct {
		Subject string
		Message string
	}

	email.Subject = "Activate your account"
	resetLink := "http://localhost:3000/auth/signup?link=" + string(token)
	email.Message = strings.ReplaceAll(emails.PasswordResetTemplate, "{{RESET_LINK}}", resetLink)

	mailErr := helpers.SendMail(user.Email, email.Subject, email.Message)
	if mailErr != nil {
		format_errors.InternalServerError(c)
		return
	}

}

// Signup function is used to create a user or signup a user
func Signup(c *gin.Context) {

	var userInput struct {
		Password string `json:"password" binding:"required,min=6"`
	}

	query := c.Query("link")

	var passwordReset userModel.PasswordReset
	initializers.DB.Where("token = ?", query).Where("expires_at > ?", time.Now()).First(&passwordReset)

	if passwordReset.ID == 0 {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid token",
			},
		)

		return
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
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

	// Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			},
		)

		return
	}

	result := initializers.DB.Model(&userModel.User{}).Where("email = ?", passwordReset.Email).Update(
		"password", string(hashPassword),
	)

	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"message": "Successfully Reset Password",
		},
	)
}

// Login function is used to log in a user
func Login(c *gin.Context) {
	// Get the email and password from the request
	var userInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if c.ShouldBindJSON(&userInput) != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			},
		)

		return
	}

	// Find the user by email
	var user userModel.User
	initializers.DB.First(&user, "email = ?", userInput.Email)

	if user.ID == 0 {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			},
		)

		return
	}

	// Compare the password with user hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			},
		)

		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)

	// Sign in and get the complete encoded token as a string using the .env secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to create token",
			},
		)
		return
	}

	// Set expiry time and send the token back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

// Logout function is used to log out a user
func Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("Authorization", "", 0, "", "", false, true)

	c.JSON(
		http.StatusOK, gin.H{
			"successMessage": "Logout successful",
		},
	)
}
