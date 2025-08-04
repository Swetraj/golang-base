package validations

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBind(obj); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.JSON(
				http.StatusUnprocessableEntity, gin.H{
					"validations": FormatValidationErrors(ve),
				},
			)
			return false
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}
