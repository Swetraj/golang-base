package helpers

import (
	"github.com/Swetraj/golang-base/api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAuthUser returns the authenticated user details from the Gin context
func GetAuthUser(c *gin.Context) *middleware.AuthUser {
	authUser, exists := c.Get("authUser")

	if !exists {
		c.JSON(
			http.StatusUnauthorized, gin.H{
				"error": "Unable to get authorized user",
			},
		)
		return nil
	}

	if user, ok := authUser.(middleware.AuthUser); ok {
		return &user
	}

	return nil
}
