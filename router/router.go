package router

import (
	"Golang-Gin-Gonic/authentication"
	"Golang-Gin-Gonic/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	router := r.Group("", authentication.BasicAuth).Group("/gin/api")

	bookRouter(router)
	router.GET("/ping", handler.HealthCheckHandler)

	return r
}

func bookRouter(r *gin.RouterGroup) {
	r.GET("books", handler.GetBooksHandler)
	r.GET("books/:bookId", handler.GetBookHandler)
	r.POST("books", handler.AddBookHandler)
	r.PUT("books/:bookId", handler.UpdateBook)
	r.DELETE("books/:bookId", handler.DeleteBook)
}
