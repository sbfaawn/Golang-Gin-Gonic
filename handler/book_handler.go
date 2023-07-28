package handler

import (
	"Golang-Gin-Gonic/service"

	"github.com/gin-gonic/gin"
)

func GetBookHandler(c *gin.Context) {
	bookId := c.Param("bookId")

	book, err := service.GetBookById(bookId)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed",
			"data":    "",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"data":    book,
		})
	}

}

func GetBooksHandler(c *gin.Context) {

}
