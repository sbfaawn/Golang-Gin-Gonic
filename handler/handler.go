package handler

import "github.com/gin-gonic/gin"

func HealthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Up",
	})
}
