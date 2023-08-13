package handler

import (
	"Golang-Gin-Gonic/dto/request"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var credential request.CredentialsRequest
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}
}
