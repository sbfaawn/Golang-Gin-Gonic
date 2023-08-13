package authentication

import (
	"Golang-Gin-Gonic/dto/response"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("x2wf9i55YRaZpemLawE6")

func JsonWebTokenAuth(ctx *gin.Context) {
	token, err := ctx.Cookie("jwt-token")
	if err != nil {
		if err == http.ErrNoCookie {
			ctx.JSON(401, response.BaseResponse{
				Message: "Unathorized - Token is missing",
				Data:    "",
				Error:   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.JSON(400, response.BaseResponse{
			Message: "Bad Request",
			Data:    "",
			Error:   err.Error(),
		})
		ctx.Abort()
		return
	}

	claims := Claims{}

	tkn, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(401, response.BaseResponse{
				Message: "Unathorized - Signature Invalid",
				Data:    "",
				Error:   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.JSON(400, response.BaseResponse{
			Message: "Bad Request",
			Data:    "",
			Error:   err.Error(),
		})
		ctx.Abort()
		return
	}
	if !tkn.Valid {
		ctx.JSON(401, response.BaseResponse{
			Message: "Unathorized - Token Invalid",
			Data:    "",
			Error:   err.Error(),
		})
		ctx.Abort()
		return
	}
}

func GenerateJWTToken(username string, expirationTime time.Time) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("error - error when creating token")
	}

	return tokenString, nil
}

func RefreshExpirationToken(tokenStr string, expirationTime time.Time) (string, error) {
	claims := Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}

		return "", err
	}
	if !tkn.Valid {
		return "", errors.New("token not valid")
	}

	if time.Until(claims.ExpiresAt.Time) > 3*time.Minute {
		return "", errors.New("token can still be used")
	}

	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
