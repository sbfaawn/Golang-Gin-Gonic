package service

import (
	authentication "Golang-Gin-Gonic/authentication/hashing"
	"Golang-Gin-Gonic/dto/request"
	"Golang-Gin-Gonic/model"
	repository "Golang-Gin-Gonic/repository/sql"
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

	// send email

	return credential, nil
}

func Login(ctx *gin.Context, credential model.Credential) error {
	result, err := repository.FindCredentialByUsername(ctx, credential.Username)

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

func AccountVerification(ctx *gin.Context, email string) error {
	err := repository.UpdateVerifiedByEmail(ctx, email)

	if err != nil {
		return err
	}

	return nil
}

func ChangePassword(ctx *gin.Context, resetPass request.ResetPasswordRequest) error {
	credential, err := repository.FindCredentialByUsername(ctx, resetPass.Username)

	if err != nil {
		return err
	}

	if authentication.IsHashedPasswordMatch(credential.Password, resetPass.NewPassword) {
		return errors.New("new password can be same with old password")
	}

	hashedPass, err := authentication.HashPassword(resetPass.NewPassword)

	if err != nil {
		return err
	}

	err = repository.UpdatePasswordByUsername(ctx, resetPass.Username, hashedPass)

	if err != nil {
		return err
	}

	return nil
}
