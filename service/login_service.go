package service

import (
	authentication "Golang-Gin-Gonic/authentication/hashing"
	"Golang-Gin-Gonic/model"
	credentialRepository "Golang-Gin-Gonic/repository/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, credential model.Credential) (model.Credential, error) {
	hashed, err := authentication.HashPassword(credential.Password)

	if err != nil {
		return credential, err
	}

	credential.Password = hashed

	err = credentialRepository.InsertCredential(ctx, credential)

	if err != nil {
		return credential, err
	}

	return credential, nil
}

func Login(ctx *gin.Context, credential model.Credential) error {
	result, err := credentialRepository.FindCredentialByUsername(ctx, credential.Username)

	if !result.IsVerified {
		return errors.New("account is not verified, please check email and do verification before login")
	}

	if err != nil {
		return err
	}

	if !authentication.IsHashedPasswordMatch(result.Password, credential.Password) {
		return errors.New("password is not correct")
	}

	return nil
}
