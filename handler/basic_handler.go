package handler

import (
	"Golang-Gin-Gonic/dto/response"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(200, response.BaseResponse{
		Message: "Up",
		Data:    "",
		Error:   "",
	})
}

func NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(404, response.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "404 Endpoint not found",
	})
}

func NoMethodAllowed(ctx *gin.Context) {
	ctx.JSON(400, response.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "No Method Allowed",
	})
}
