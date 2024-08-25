package handler

import (
	"api-pustaka/book"
	"api-pustaka/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.BookService
}

func NewBookHandler(bookService book.BookService) *bookHandler  {
	return &bookHandler{bookService}
}

func (h* bookHandler) GetBooks(c *gin.Context)  {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := helper.ConvertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context)  {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	bookResponse := helper.ConvertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context)   {
	var bookRequest book.BookRequest

	// Validation Error
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// Logic
	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"data": helper.ConvertToBookResponse(book),
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context)   {
	var bookRequest book.BookRequest

	// Validation Error
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// Logic
	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"data": helper.ConvertToBookResponse(book),
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context)   {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// Logic
	_, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"data": "Data successfully delated",
	})
}