package repository

import (
	"Golang-Gin-Gonic/configuration"
	"Golang-Gin-Gonic/model"
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

var conditionSoftDelete = "deleted_at IS null"

func init() {
	db = configuration.DB
}

func FindBookById(ctx context.Context, id uint) (model.Book, error) {
	book := model.Book{}
	var err error

	tx := db.Begin()
	err = tx.Where(conditionSoftDelete).First(&book, id).Error

	if err != nil {
		tx.Rollback()
		return model.Book{}, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return model.Book{}, err
	}

	return book, nil
}

func FindBooks(ctx context.Context) ([]model.Book, error) {
	books := []model.Book{}
	var err error

	tx := db.Begin()
	err = tx.Where(conditionSoftDelete).Find(&books).Error

	if err != nil {
		tx.Rollback()
		return books, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return books, err
	}

	return books, nil
}

func FindBookByDetail(ctx context.Context, book model.Book) (model.Book, error) {
	var err error
	result := model.Book{}

	tx := db.Begin()
	err = tx.Find(&result).Where("(author = ? AND title = ?) AND deleted_at IS NULL", book.Author, book.Title).Error

	if err != nil {
		tx.Rollback()
		return model.Book{}, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return model.Book{}, err
	}

	return result, nil
}

func AddBook(ctx context.Context, book model.Book) error {
	var err error

	tx := db.Begin()
	err = tx.Create(&book).Error

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

func UpdateBookById(ctx context.Context, id uint, book model.Book) error {
	var err error

	tx := db.Begin()

	book.Id = id
	update := tx.Model(&book).Where(conditionSoftDelete).Updates(&book)

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
		return errors.New("record with id " + strconv.FormatUint(uint64(id), 10) + " is not found")
	}

	return nil
}

func DeleteBookById(ctx context.Context, id uint) (model.Book, error) {
	var err error

	book := model.Book{
		Id: id,
		DeletedAt: gorm.DeletedAt{
			Time:  time.Now(),
			Valid: true,
		},
	}

	tx := db.Begin()
	delete := tx.Model(&book).Where("deleted_at IS null").Delete(&book)

	if delete.Error != nil {
		tx.Rollback()
		return model.Book{}, err
	}

	result := delete.Commit().WithContext(ctx)

	if result.Error != nil {
		return model.Book{}, err
	}

	if delete.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return model.Book{}, errors.New("record with id " + strconv.FormatUint(uint64(id), 10) + " is not found")
	}

	return book, nil
}
