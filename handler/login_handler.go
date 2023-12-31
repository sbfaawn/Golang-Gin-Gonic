package handler

import (
	"Golang-Gin-Gonic/authorization"
	"Golang-Gin-Gonic/dto/request"
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/service"
	"Golang-Gin-Gonic/validator"
	"time"

	"github.com/gin-gonic/gin"
)

var tokenLifetime int = 5

var jwtTokenKey string = "jwt-token"

func RegisterHandler(ctx *gin.Context) {
	var request request.CredentialsRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	if err := isLoginRequestValid(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	credential := model.Credential{
		Username:   request.Username,
		Password:   request.Password,
		Email:      request.Email,
		IsVerified: false,
	}

	result, err := service.Register(ctx, credential)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, result, nil)
}

func LoginHandler(ctx *gin.Context) {
	var request request.CredentialsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	credential := model.Credential{
		Username: request.Username,
		Password: request.Password,
	}

	err := service.Login(ctx, credential)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	expirationTime := time.Now().Add(time.Duration(tokenLifetime) * time.Minute)
	token, err := authorization.GenerateJWTToken(credential.Username, expirationTime)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	ctx.SetCookie(jwtTokenKey, token, tokenLifetime*int(time.Minute), "/", "localhost", false, true)
	generateResponse(ctx, 200, "", nil)
}

func RefreshTokenHandler(ctx *gin.Context) {
	token, _ := ctx.Cookie(jwtTokenKey)

	expirationTime := time.Now().Add(time.Duration(tokenLifetime) * time.Minute)
	newToken, err := authorization.RefreshExpirationToken(token, expirationTime)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	ctx.SetCookie(jwtTokenKey, newToken, tokenLifetime*int(time.Minute), "/", "localhost", false, true)
	generateResponse(ctx, 200, "", nil)
}

func LogoutHandler(ctx *gin.Context) {
	ctx.SetCookie(jwtTokenKey, "", -1, "/", "localhost", false, true)
	generateResponse(ctx, 200, "", nil)
}

func isLoginRequestValid(credential *request.CredentialsRequest) error {
	validate := validator.Validate

	err := validate.Struct(credential)

	if err != nil {
		return err
	}

	return nil
}

func AccountVerificationHandler(ctx *gin.Context) {
	emailParam := ctx.Param("email")

	err := service.AccountVerification(ctx, emailParam)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, "", nil)
}

func ResetPasswordHandler(ctx *gin.Context) {
	// send email
}

func ChangePasswordHandler(ctx *gin.Context) {
	var resetPass request.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&resetPass); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	err := service.ChangePassword(ctx, resetPass)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, "", nil)
}
