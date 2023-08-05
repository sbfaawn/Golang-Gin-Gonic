package handler

import (
	"Golang-Gin-Gonic/dto/request"
	"Golang-Gin-Gonic/dto/response"
	"Golang-Gin-Gonic/model"
	"Golang-Gin-Gonic/service"
	"Golang-Gin-Gonic/validator"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookHandler(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	bookIdNum, _ := strconv.ParseUint(bookId, 10, 32)

	book, err := service.GetBook(ctx, uint(bookIdNum))

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, book, nil)
}

func GetBooksHandler(ctx *gin.Context) {
	books, err := service.GetBooks(ctx)

	if err != nil {
		generateResponse(ctx, 400, "", err)
	}

	if len(books) == 0 {
		generateResponse(ctx, 400, "", errors.New("no books is stored right now"))
		return
	}

	generateResponse(ctx, 200, books, nil)

}

func AddBookHandler(ctx *gin.Context) {
	var request request.BookRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	if err := isRequestValid(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	book := model.Book{
		Title:  request.Title,
		Author: request.Author,
	}

	result, err := service.AddBook(ctx, book)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, result, nil)

}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	bookIdNum, _ := strconv.ParseUint(bookId, 10, 32)

	var request request.BookRequest
	if err := ctx.BindJSON(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	if err := isRequestValid(&request); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	book := model.Book{
		Title:  request.Title,
		Author: request.Author,
	}

	result, err := service.UpdateBook(ctx, uint(bookIdNum), book)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, result, nil)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	bookIdNum, _ := strconv.ParseUint(bookId, 10, 32)

	result, err := service.DeleteBook(ctx, uint(bookIdNum))

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, result, nil)
}

func generateResponse(ctx *gin.Context, statusCode int, data any, err error) {
	var message string

	message = "Success"
	if statusCode != 200 {
		message = "Failed"
	}

	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	ctx.JSON(statusCode, response.BaseResponse{
		Message: message,
		Data:    data,
		Error:   errorMessage,
	})
}

func isRequestValid(request *request.BookRequest) error {
	validate := validator.Validate

	err := validate.Struct(request)

	if err != nil {
		return errors.New("title & author field need to specific on request body")
	}

	return nil
}
