package service

import (
	authentication "Golang-Gin-Gonic/authentication/hashing"
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/repository"
	"errors"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, credential model.Credential) (model.Credential, error) {
	hashed, err := authentication.HashPassword(credential.Password)

	if err != nil {
		return credential, err
	}

	credential.Password = hashed

	err = repository.InsertCredential(ctx, credential)

	if err != nil {
		return credential, err
	}

	return credential, nil
}

func Login(ctx *gin.Context, credential model.Credential) error {
	result, err := repository.FindCredentialByUsername(ctx, credential.Username)

	if err != nil {
		return err
	}

	if !authentication.IsHashedPasswordMatch(result.Password, credential.Password) {
		return errors.New("password is not correct")
	}

	return nil
}
