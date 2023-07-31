package handler

import (
	"Golang-Gin-Gonic/dto/request"
	"Golang-Gin-Gonic/dto/response"
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/service"

	"github.com/gin-gonic/gin"
)

func GetBookHandler(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	book, err := service.GetBook(bookId)

	if err != nil {
		generateResponse(ctx, 400, err)
		return
	}

	generateResponse(ctx, 200, book)

}

func GetBooksHandler(ctx *gin.Context) {
	books := service.GetBooks()

	if len(books) == 0 {
		generateResponse(ctx, 400, "")
		return
	}

	generateResponse(ctx, 200, books)

}

func AddBookHandler(ctx *gin.Context) {
	var request request.BookRequest
	if err := ctx.BindJSON(&request); err != nil {
		generateResponse(ctx, 400, err)
		return
	}

	book := model.Book{
		Id:     "",
		Title:  request.Name,
		Author: request.Author,
	}

	result, err := service.AddBook(book)

	if err != nil {
		generateResponse(ctx, 400, err)
		return
	}

	generateResponse(ctx, 200, result)

}

func generateResponse(ctx *gin.Context, statusCode int, data any) {
	var message string

	message = "Success"
	if statusCode != 200 {
		message = "Failed"
	}

	ctx.JSON(statusCode, response.BaseResponse{
		Message: message,
		Data:    data,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	var request request.BookRequest
	if err := ctx.BindJSON(&request); err != nil {
		generateResponse(ctx, 400, err)
		return
	}

	book := model.Book{
		Id:     "",
		Title:  request.Name,
		Author: request.Author,
	}

	result, err := service.UpdateBook(bookId, book)

	if err != nil {
		generateResponse(ctx, 400, err)
		return
	}

	generateResponse(ctx, 200, result)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	result, err := service.DeleteBook(bookId)

	if err != nil {
		generateResponse(ctx, 400, err)
		return
	}

	generateResponse(ctx, 200, result)
}
