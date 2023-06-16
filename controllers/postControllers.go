package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyper-dot/go-crud-api/initializers"
	"github.com/hyper-dot/go-crud-api/models"
)

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
