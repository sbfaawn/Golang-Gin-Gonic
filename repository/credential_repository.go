package repository

import (
	"Golang-Gin-Gonic/model"
	"errors"

	"github.com/gin-gonic/gin"
)

func InsertCredential(ctx *gin.Context, credential model.Credential) error {
	var err error

	tx := db.Begin()
	err = tx.Create(&credential).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func FindCredentialByUsernamePassword(ctx *gin.Context, credential model.Credential) error {
	var err error

	tx := db.Begin()
	err = tx.Where("deleted_at IS null").First(&credential, "username = ? AND password = ?", credential.Username, credential.Password).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func UpdatePasswordByUsername(ctx *gin.Context, username string, newPassword string) error {
	var err error

	tx := db.Begin()
	update := tx.Model(&model.Credential{}).Where("username = ? AND deleted_at IS null", username).Updates(map[string]any{
		"password": newPassword,
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	result := update.Commit().WithContext(ctx)

	if result.Error != nil {
		tx.Rollback()
		return err
	}

	if update.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return errors.New("record with username " + username + " is not found")
	}

	return nil
}
