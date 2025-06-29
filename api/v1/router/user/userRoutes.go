package userRoutes

import (
	controllers2 "github.com/Swetraj/golang-base/api/v1/controllers/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	println("register user routes")
	userRouter := rg.Group("/users")
	{
		userRouter.GET("/", controllers2.GetUsers)
		userRouter.GET("/:id/edit", controllers2.EditUser)
		userRouter.PUT("/:id/update", controllers2.UpdateUser)
		userRouter.DELETE("/:id/delete", controllers2.DeleteUser)
		userRouter.GET("/all-trash", controllers2.GetTrashedUsers)
	}
}
