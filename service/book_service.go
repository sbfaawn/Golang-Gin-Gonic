package service

import (
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/repository"
	"errors"
)

func GetBooks() []model.Book {
	return repository.FindBooks()
}

func GetBook(id string) (model.Book, error) {
	book, err := repository.FindBookById(id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func AddBook(book model.Book) (model.Book, error) {
	_, err := repository.FindBookByDetail(book)

	if err == nil {
		return model.Book{}, errors.New("book with particular detail is already stored")
	}

	result, err := repository.AddBook(book)

	if err != nil {
		return model.Book{}, err
	}

	return result, nil
}

func UpdateBook(id string, book model.Book) (model.Book, error) {
	result, err := repository.UpdateBookById(id, book)

	if err != nil {
		return model.Book{}, err
	}

	return result, nil
}

func DeleteBook(id string) (model.Book, error) {
	book, err := repository.DeleteBookById(id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}
