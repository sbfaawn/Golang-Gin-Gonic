package service

import (
	"Golang-Gin-Gonic/model"
	bookRepository "Golang-Gin-Gonic/repository/sql"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) ([]model.Book, error) {
	books, err := bookRepository.FindBooks(ctx)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func GetBook(ctx *gin.Context, id uint) (model.Book, error) {
	book, err := bookRepository.FindBookById(ctx, id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func AddBook(ctx *gin.Context, book model.Book) (model.Book, error) {
	err := bookRepository.AddBook(ctx, book)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func UpdateBook(ctx *gin.Context, id uint, book model.Book) (model.Book, error) {
	err := bookRepository.UpdateBookById(ctx, id, book)

	if err != nil {
		return model.Book{}, err
	}

	book.Id = id
	return book, nil
}

func DeleteBook(ctx *gin.Context, id uint) (model.Book, error) {
	book, err := bookRepository.DeleteBookById(ctx, id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}
