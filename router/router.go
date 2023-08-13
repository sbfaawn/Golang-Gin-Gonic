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
	r.NoMethod(handler.NoMethodAllowed)

	router := r.Group("", authentication.BasicAuth).Group("/api")
	// check jwt token on every route except login, register, ping
	baseRouter(router)

	authRouter := router.Group("/auth")
	loginRouter(authRouter)
	productRouter := router.Group("/products", authentication.JsonWebTokenAuth)
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
	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)
	r.GET("/refresh", authentication.JsonWebTokenAuth, handler.RefreshTokenHandler)
	r.GET("/logout", handler.LogoutHandler)
}

func baseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.HealthCheckHandler)
}
