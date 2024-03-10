package controllers

import (
	"crud-api/initializers"
	"crud-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost (c *gin.Context) {
	// get data off req body

	var data struct{
		Body string 
		Title string
	}

	c.Bind(&data)
	post := models.Post{Title: data.Title,Body: data.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		return
	}
	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context){

	// get the posts
	var posts []models.Post

	initializers.DB.Find(&posts)

	// response with them
	c.JSON(200, gin.H{
		"posts":posts,
	})
}
func GetSinglePost(c *gin.Context){
	// get id from url
	 id:= c.Param("id")
	// get the post
	var post models.Post

	initializers.DB.First(&post, id)

	// response with them
	c.JSON(200, gin.H{
		"post":post,
	})
}

func UpdatePost (c *gin.Context){
	// get the id from the url
	id:= c.Param("id")

	// get the data from req body
	var data struct{
		Body string 
		Title string
	}
	c.Bind(&data)

	// find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: data.Title,
		Body: data.Body, 
	})

	// respond wiht it
	c.JSON(200, gin.H{
		"post":post,
	})
}

// func DeletePost (c *gin.Context){
// 	// get the id from the url
// 	id:= c.Param("id") 

// 	// delete the post
	
// 	initializers.DB.Delete(&models.Post{},id)

// 	// respond with it
// 	c.JSON(200, gin.H{
// 		"message":"post deleted",
// 	})
// }

func DeletePost(c *gin.Context) {
	// Get the id from the url
	id := c.Param("id")

	// Check if the post exists
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// If the post is not found, respond with an error message
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Post not found",
			})
		} else {
			// If there was an error other than not found, respond with an error message
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error checking post existence",
				"error":   err.Error(),
			})
		}
		return
	}

	// If the post exists, proceed to delete it
	if err := initializers.DB.Delete(&post).Error; err != nil {
		// If there was an error during deletion, respond with an error message
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting post",
			"error":   err.Error(),
		})
		return
	}

	// If the deletion was successful, respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted",
	})
}