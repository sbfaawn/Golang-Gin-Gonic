package repository

import (
	"Golang-Gin-Gonic/model"
	"errors"
	"strconv"
)

func FindBookById(id string) (model.Book, error) {
	for _, book := range model.Books {
		if book.Id == id {
			return book, nil
		}
	}

	return model.Book{}, errors.New("book with id" + id + "is not found")
}

func FindBooks() []model.Book {
	return model.Books
}

func FindBookByDetail(bookSearch model.Book) (model.Book, error) {
	for _, book := range model.Books {
		if book.Title == bookSearch.Title && book.Author == bookSearch.Author {
			return book, nil
		}
	}

	return model.Book{}, errors.New("book with Author " + bookSearch.Author + " and title " + bookSearch.Title + " is not found")
}

func AddBook(book model.Book) (model.Book, error) {
	book.Id = strconv.Itoa(len(model.Books) + 1)
	model.Books = append(model.Books, book)

	if model.Books[len(model.Books)-1] == book {
		return book, nil
	}

	return model.Book{}, errors.New("failed to add book to collection")
}
