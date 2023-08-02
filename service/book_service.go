package service

import (
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/repository"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) ([]model.Book, error) {
	books, err := repository.FindBooks(ctx)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func GetBook(ctx *gin.Context, id uint) (model.Book, error) {
	book, err := repository.FindBookById(ctx, id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func AddBook(ctx *gin.Context, book model.Book) (model.Book, error) {
	_, err := repository.FindBookByDetail(ctx, book)
	fmt.Println("1")

	if err != nil {
		return model.Book{}, errors.New("book with particular detail is already stored")
	}

	fmt.Println("2")
	result, err := repository.AddBook(ctx, book)

	fmt.Println("3")
	if err != nil {
		return model.Book{}, err
	}

	fmt.Println("4")
	return result, nil
}

func UpdateBook(ctx *gin.Context, id uint, book model.Book) (model.Book, error) {
	result, err := repository.UpdateBookById(ctx, id, book)

	if err != nil {
		return model.Book{}, err
	}

	return result, nil
}

func DeleteBook(ctx *gin.Context, id string) (model.Book, error) {
	book, err := repository.DeleteBookById(ctx, id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}
