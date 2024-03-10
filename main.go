package main

import (
	"crud-api/controllers"
	"crud-api/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
 initializers.LoadEnvVariables()
 initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run() 
}