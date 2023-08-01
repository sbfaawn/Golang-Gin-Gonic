package authentication

import (
	"Golang-Gin-Gonic/dto/response"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	username, password, isOk := ctx.Request.BasicAuth()

	if !(isOk && username == "dhika" && password == "password1234") {
		ctx.JSON(401, response.BaseResponse{
			Data:    nil,
			Message: "unauthorized user, username & password need to access this resource",
		})
		ctx.Abort()
		ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}
