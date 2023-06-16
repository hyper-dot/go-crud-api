package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hyper-dot/go-crud-api/controllers"
	"github.com/hyper-dot/go-crud-api/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("hello123")
	r := gin.Default()
	r.POST("/", controllers.CreatePost)

	r.Run()
}
