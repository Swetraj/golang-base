package userController

import (
	"github.com/Swetraj/golang-base/db/initializers"
	formatErrors "github.com/Swetraj/golang-base/internal/format-errors"
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EditUser function is used to find a user by id
func EditUser(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the user
	var user userModel.User
	result := initializers.DB.First(&user, id)

	if err := result.Error; err != nil {
		formatErrors.RecordNotFound(c, err)
		return
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}
