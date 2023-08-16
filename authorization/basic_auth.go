package authorization

import (
	"Golang-Gin-Gonic/dto/response"
	properties_reader "Golang-Gin-Gonic/properties/reader"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	bac := properties_reader.Config.Auth.ConfigBasicAuth
	username, password, isOk := ctx.Request.BasicAuth()

	if !(isOk && username == bac.Username && password == bac.Password) {
		ctx.JSON(401, response.BaseResponse{
			Data:    nil,
			Message: "unauthorized user, username & password need to access this resource",
		})
		ctx.Abort()
		ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}
