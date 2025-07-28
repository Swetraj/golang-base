package main

import (
	"fmt"
	"github.com/Swetraj/golang-base/api/middleware"
	"github.com/Swetraj/golang-base/api/v1/router"
	"github.com/Swetraj/golang-base/config"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("Hello auth")
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	router.GetRoute(r)
	err := r.Run()
	if err != nil {
		return
	}
}
