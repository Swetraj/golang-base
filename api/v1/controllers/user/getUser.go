package userController

import (
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/dto"
	formatErrors "github.com/Swetraj/golang-base/internal/format-errors"
	userModel "github.com/Swetraj/golang-base/internal/models/user"
	"github.com/Swetraj/golang-base/internal/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	// Get all the users
	var users []userModel.User

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	perPageStr := c.DefaultQuery("perPage", "5")
	perPage, _ := strconv.Atoi(perPageStr)

	result, err := pagination.Paginate(initializers.DB, page, perPage, nil, &users)
	if err != nil {
		formatErrors.InternalServerError(c)
		return
	}

	userDto, err := dto.UsersToUserResponseDTOs(users)

	if err != nil {
		formatErrors.InternalServerError(c)
		return
	}

	result.Data = userDto

	// Return the users
	c.JSON(http.StatusOK, gin.H{
		"data": result.Data,
		"meta": result.Meta,
	})
}
