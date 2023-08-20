package router

import (
	"Golang-Gin-Gonic/authorization"
	"Golang-Gin-Gonic/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.NoRoute(handler.NoRouteHandler)
	r.NoMethod(handler.NoMethodAllowed)

	router := r.Group("", authorization.BasicAuth).Group("/api")
	// check jwt token on every route except login, register, ping
	baseRouter(router)

	authRouter := router.Group("/auth")
	loginRouter(authRouter)
	productRouter := router.Group("/products", authorization.JsonWebTokenAuth)
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
	r.GET("/refresh", authorization.JsonWebTokenAuth, handler.RefreshTokenHandler)
	r.GET("/logout", handler.LogoutHandler)
	r.POST("/password/change", handler.ChangePasswordHandler)
	r.POST("/password/reset", handler.ResetPasswordHandler)
	r.GET("/verification/:email", handler.AccountVerificationHandler)
}

func baseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.HealthCheckHandler)
}
