package controllers

import (
	"net/http"

	"github.com/edutjie/go-crud/initializers"
	"github.com/edutjie/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get the user
	user := c.MustGet("user").(models.User)

	// get the request body
	var body struct {
		Title string
		Body  string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding error: " + err.Error()})
		return
	}

	// create a new post
	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		UserID: user.ID,
	}

	// save the post to the database
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Creating post error: " + result.Error.Error()})
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Finding posts error: " + result.Error.Error()})
		return
	}

	// return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsUserIndex(c *gin.Context) {
	// get the user
	user := c.MustGet("user").(models.User)

	// get all posts
	var posts []models.Post
	result := initializers.DB.Where("user_id = ?", user.ID).Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Finding posts error: " + result.Error.Error()})
		return
	}

	// return the posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id
	id := c.Param("id")

	// get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get user
	user := c.MustGet("user").(models.User)

	// get id
	id := c.Param("id")

	// get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// check if the user is the owner of the post
	if post.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not the owner of this post",
		})
		return
	}

	// get the request body
	var body struct {
		Title string
		Body  string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Binding error: " + err.Error()})
		return
	}

	// first way: downside is that it will update all the fields even if they are not provided
	// // update the post
	// post.Title = body.Title
	// post.Body = body.Body

	// // save the post
	// result = initializers.DB.Save(&post)

	// if result.Error != nil {
	// 	c.Status(500)
	// 	return
	// }

	// second way
	// update the post
	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Updating post error: " + result.Error.Error()})
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get user
	user := c.MustGet("user").(models.User)

	// get id
	id := c.Param("id")

	// get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// check if the user is the owner of the post
	if post.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not the owner of this post",
		})
		return
	}

	// delete the post
	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deleting post error: " + result.Error.Error()})
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}
