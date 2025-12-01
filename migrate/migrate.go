package main

import (
	"learn_gin/initializers"
	"learn_gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectWithDb()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.Comment{})
}
