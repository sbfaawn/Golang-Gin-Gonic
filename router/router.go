package router

import (
	"Golang-Gin-Gonic/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("/gin/api")

	bookRouter(router)

	router.GET("/ping", handler.HealthCheckHandler)

	return r
}

func bookRouter(r *gin.RouterGroup) {
	// r.GET("users", handler.)
	r.GET("book/:bookId", handler.GetBookHandler)
}
