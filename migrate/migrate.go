package main

import (
	"github.com/edutjie/go-crud/initializers"
	"github.com/edutjie/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
}
