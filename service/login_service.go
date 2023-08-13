package service

import (
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/repository"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, credential model.Credential) (model.Credential, error) {
	err := repository.InsertCredential(ctx, credential)

	if err != nil {
		return credential, err
	}

	return credential, nil
}

func Login(ctx *gin.Context, credential model.Credential) error {
	err := repository.FindCredentialByUsernamePassword(ctx, credential)

	if err != nil {
		return err
	}

	return nil
}
