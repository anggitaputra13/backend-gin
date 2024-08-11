package controllers

import (
	"net/http"

	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	format_errors "github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/format-errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/pagination"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	// Get data from request
	var bookInput models.BookRequest
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Title unique validation
	if validations.IsUniqueValue("books", "title", bookInput.Title, 0) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Title": "The title is already exist!",
			},
		})

		return
	}

	// Create the book
	book := models.Book{
		Title:    bookInput.Title,
		Category: bookInput.Category,
		Price:    bookInput.Price,
		Qty:      bookInput.Qty,
	}

	result := initializers.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot create book",
		})

		return
	}

	// Return the book
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

// GetAll Book
func ListBook(c *gin.Context) {
	var allBook []models.Book

	var filter models.BookFilter
	if err := c.BindQuery(&filter); err != nil {
		format_errors.InternalServerError(c)
		return
	}

	result, err := pagination.Paginate(initializers.DB, filter.Page, filter.Limit, nil, &allBook)

	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the books
	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}

// get the book by ID
func GetBook(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the book
	var book models.Book
	result := initializers.DB.First(&book, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

// UpdateBook updates a book
func UpdateBook(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	var bookInput models.BookRequest
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Find the book by ID
	var book models.Book
	result := initializers.DB.First(&book, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Name unique validation
	if validations.IsUniqueValue("books", "title", bookInput.Title, book.ID) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Title": "The title is already exist!",
			},
		})

		return
	}

	updateBook := models.Book{
		Title:    bookInput.Title,
		Category: bookInput.Category,
		Price:    bookInput.Price,
		Qty:      bookInput.Qty,
	}

	// Update the book record
	result = initializers.DB.Model(&book).Updates(updateBook)
	if err := result.Error; err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the book
	c.JSON(http.StatusOK, gin.H{
		"book": updateBook,
	})
}

// Soft deletes a book by id
func DeleteBook(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")
	var book models.Book

	// Find the book
	result := initializers.DB.First(&book, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the post
	initializers.DB.Delete(&book)

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The book has been deleted successfully",
	})
}

// Permanent deletes a book by id
func DeleteBookPermanent(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")

	// Delete the post
	result := initializers.DB.Unscoped().Delete(&models.Book{}, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The book has been deleted permanently",
	})
}
