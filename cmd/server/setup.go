package main

import (
	"github.com/Swetraj/golang-base/api/middleware"
	handler2 "github.com/Swetraj/golang-base/api/v1/handler"
	"github.com/Swetraj/golang-base/db/initializers"
	"github.com/Swetraj/golang-base/internal/repository"
	services2 "github.com/Swetraj/golang-base/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupApp() *gin.Engine {
	db := initializers.DB
	repositories := repository.NewRepositories(db)
	services := services2.NewServices(repositories)
	handler := handler2.NewHandler(services)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	handler.RegisterRoutes(api)
	return r
}
