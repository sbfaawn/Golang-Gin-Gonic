package repository

import (
	"Golang-Gin-Gonic/configuration"
	"Golang-Gin-Gonic/model"
	"context"
	"time"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = configuration.DB
}

func FindBookById(ctx context.Context, id uint) (model.Book, error) {
	book := model.Book{}
	var err error

	tx := db.Begin()
	err = tx.Where("deleted_at IS null").First(&book, id).Error

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
	err = tx.Where("deleted_at IS null").Find(&books).Error

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

func AddBook(ctx context.Context, book model.Book) (model.Book, error) {
	var err error

	tx := db.Begin()
	err = tx.Create(&book).Error

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

func UpdateBookById(ctx context.Context, id uint, book model.Book) (model.Book, error) {
	var err error
	book.Id = id

	fmt.Println("new book : ", book)
	tx := db.Begin()
	err = tx.Model(model.Book{}).Where("id = ? AND deleted_at IS NULL", id).Updates(map[string]any{
		"author":book.Author,
		"title":book.Title,
	}).Error

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

func DeleteBookById(ctx context.Context, id string) (model.Book, error) {
	var err error
	var book model.Book

	tx := db.Begin()
	err = tx.Model(model.Book{}).Where("id = ? AND deleted_at IS null", id).Update("deleted_at", time.Now()).Error

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
