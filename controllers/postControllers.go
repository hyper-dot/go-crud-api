package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyper-dot/go-crud-api/initializers"
	"github.com/hyper-dot/go-crud-api/models"
)

// Create Posts
func CreatePost(c *gin.Context) {
	// Get body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	//Create Post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Error",
		})
	}
	// Return it

	c.JSON(200, gin.H{
		"Post": post,
	})
}

// Find all posts
func GetAllPost(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}

// Get a single Post

func GetPost(c *gin.Context) {
	// Get Id from url
	id := c.Param("id")
	// Find the POST
	var post models.Post
	initializers.DB.First(&post, id)
	// Return It
	c.JSON(200, gin.H{
		"Post": post,
	})
}

// Update post

func UpdatePost(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	// get the body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Find the post
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Return It
	c.JSON(200, gin.H{
		"Updated Post": post,
	})
}

// Detete Post
func DeletePost(c *gin.Context) {
	// Get Id
	id := c.Param("id")
	// Find and delete post
	post := models.Post{}

	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"message": "Deleted Post Successfully!!!",
	})

}
