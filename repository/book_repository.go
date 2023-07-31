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

func UpdateBookById(id string, bookSearch model.Book) (model.Book, error) {
	for index, book := range model.Books {
		if book.Id == id {
			bookPointer := &model.Books[index]
			bookPointer.Author = bookSearch.Author
			bookPointer.Title = bookSearch.Title
			return bookSearch, nil
		}
	}

	return model.Book{}, errors.New("cant find book with id " + id)
}

func DeleteBookById(id string) (model.Book, error) {
	for index, book := range model.Books {
		if book.Id == id {
			model.Books = append(model.Books[:index], model.Books[index+1])
			return book, nil
		}
	}

	return model.Book{}, errors.New("cant find book with id " + id)
}
