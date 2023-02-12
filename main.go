package main

import (
	"github.com/edutjie/go-crud/controllers"
	"github.com/edutjie/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// auth routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	// posts routes
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
