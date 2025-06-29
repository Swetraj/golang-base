package main

import (
	"fmt"
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
	router.GetRoute(r)

	r.Run()
}
