package userController

import (
	"github.com/Swetraj/golang-base/db/initializers"
	formatErrors "github.com/Swetraj/golang-base/internal/format-errors"
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"github.com/Swetraj/golang-base/internal/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteUser function is used to delete a user by id
func DeleteUser(c *gin.Context) {
	// Get the id from the url
	id := c.Param("id")
	var user userModel.User

	result := initializers.DB.First(&user, id)
	if err := result.Error; err != nil {
		formatErrors.RecordNotFound(c, err)
		return
	}

	// Delete the user
	initializers.DB.Delete(&user)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The user has been deleted successfully",
	})
}

// GetTrashedUsers function is used to get all the trashed user
func GetTrashedUsers(c *gin.Context) {
	// Get the users
	var users []userModel.User

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(initializers.DB.Unscoped().Where("deleted_at IS NOT NULL"), page, perPage, nil, &users)
	if err != nil {
		formatErrors.InternalServerError(c)
		return
	}

	//result := initializers.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	//if err := result.Error; err != nil {
	//	format_errors.InternalServerError(c)
	//	return
	//}

	// Return the users
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
