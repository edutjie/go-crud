package main

import (
	"github.com/edutjie/go-crud/controllers"
	"github.com/edutjie/go-crud/initializers"
	"github.com/edutjie/go-crud/middleware"
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
	r.GET("/auth/logout", controllers.Logout)
	r.GET("/auth/validate", middleware.RequireAuth, controllers.Validate)

	// posts routes
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/user", middleware.RequireAuth, controllers.PostsUserIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.PostsUpdate)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
