package controllers

import (
	"learn_gin/initializers"
	"learn_gin/models"
	"time"
	"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"fmt"
)

func PostsCreate(c *gin.Context) {
	fmt.Println("here")
	newPost :=models.Post{Title:"About GIN", Body:"Gin is easy but syntax is from alien scripts", Tags: pq.StringArray{"Edu"}, CreatedAt:time.Now()}
	result:=initializers.DB.Create(&newPost)
	fmt.Println("but here?")

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		
		"post":newPost,
	})
}