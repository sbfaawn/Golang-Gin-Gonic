package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Gin Gonic Server is Running")

	r := gin.Default()
	router := r.Group("/gin/api")
	router.GET("/ping", handler)
	r.Run(":8085") // listen and serve on 0.0.0.0:8080
}

func handler(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}