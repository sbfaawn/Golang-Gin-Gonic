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

func FindCredentialByUsername(ctx *gin.Context, username string) (model.Credential, error) {
	var err error
	var credential model.Credential

	tx := db.Begin()
	err = tx.First(&credential, "username = ?", username).Error

	if err != nil {
		tx.Rollback()
		return credential, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return credential, err
	}

	return credential, nil
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

func UpdateVerifiedByEmail(ctx *gin.Context, email string) error {
	var err error

	tx := db.Begin()
	update := tx.Model(&model.Credential{}).Where("email = ? AND deleted_at IS null", email).Updates(map[string]any{
		"verified": true,
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
		return errors.New("record with email " + email + " is not found")
	}

	return nil
}
