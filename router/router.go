package router

import (
	"Golang-Gin-Gonic/authentication"
	"Golang-Gin-Gonic/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.NoRoute(handler.NoRouteHandler)
	r.NoRoute(handler.NoMethodAllowed)

	router := r.Group("", authentication.BasicAuth).Group("/api")
	baseRouter(router)
	loginRouter(router)

	productRouter := router.Group("/products")
	bookRouter(productRouter)

	return r
}

func bookRouter(r *gin.RouterGroup) {
	r.GET("books", handler.GetBooksHandler)
	r.GET("books/:bookId", handler.GetBookHandler)
	r.POST("books", handler.AddBookHandler)
	r.PUT("books/:bookId", handler.UpdateBook)
	r.DELETE("books/:bookId", handler.DeleteBook)
}

func loginRouter(r *gin.RouterGroup) {
	r.POST("/register")
	r.POST("/login")
	r.POST("/refresh")
	r.POST("/logout")
}

func baseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.HealthCheckHandler)
}
