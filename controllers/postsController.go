package controllers

import (
	"github.com/edutjie/go-crud/initializers"
	"github.com/edutjie/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get the request body
	var body struct {
		Title string
		Body  string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(400)
		return
	}

	// create a new post
	post := models.Post{Title: body.Title, Body: body.Body}

	// save the post to the database
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(500)
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
		c.Status(500)
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
		c.Status(404)
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id
	id := c.Param("id")

	// get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	// get the request body
	var body struct {
		Title string
		Body  string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(400)
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
		c.Status(500)
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsDelete(c *gin.Context) {
	// get id
	id := c.Param("id")

	// get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(404)
		return
	}

	// delete the post
	result = initializers.DB.Delete(&post)

	if result.Error != nil {
		c.Status(500)
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}