package main

import (
	"github.com/Swetraj/golang-base/config"
	"github.com/Swetraj/golang-base/db/initializers"
)

func init() {
	config.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := SetupApp()
	err := r.Run()
	if err != nil {
		return
	}
}
