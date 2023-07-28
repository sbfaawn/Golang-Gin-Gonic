package service

import (
	"Golang-Gin-Gonic/model"
	"errors"
)

func GetBooks() []model.Book {
	return model.Books
}

func GetBookById(id string) (model.Book, error) {
	for _, book := range model.Books {
		if book.Id == id {
			return book, nil
		}
	}

	return model.Book{}, errors.New("Book with Id" + id + "is not found")
}
