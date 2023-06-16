package main

import (
	"github.com/hyper-dot/go-crud-api/initializers"
	"github.com/hyper-dot/go-crud-api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
