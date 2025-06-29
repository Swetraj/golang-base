package userController

import (
	"errors"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/dto"
	formatErrors "github.com/Swetraj/golang-base/internal/format-errors"
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"github.com/Swetraj/golang-base/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// UpdateUser function is used to update a user
func UpdateUser(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the name, email and password from request
	var userInput struct {
		Name  string `json:"name" binding:"required,min=2,max=50"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the user by id
	var user userModel.User
	result := initializers.DB.First(&user, id)

	if err := result.Error; err != nil {
		formatErrors.RecordNotFound(c, err)
		return
	}

	// Email unique validation
	if user.Email != userInput.Email && validations.IsUniqueValue("users", "email", userInput.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"validations": map[string]interface{}{
				"Email": "The email is already exist!",
			},
		})
		return
	}

	// Prepare data to update
	updateUser := userModel.User{
		Email: userInput.Email,
	}

	// Update the user
	result = initializers.DB.Model(&user).Updates(&updateUser)

	userDto := dto.SingleUserToUserResponseDTO(user)

	if result.Error != nil {
		formatErrors.InternalServerError(c)
		return
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"data": userDto,
	})
}
