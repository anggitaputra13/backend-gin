package controllers

import (
	"net/http"
	"time"

	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/db/initializers"
	format_errors "github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/format-errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/pagination"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// CreateOrder creates a order
func CreateOrder(c *gin.Context) {
	// Get input from request
	var orderInput models.OrderRequest

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	//populate total price
	var totalPrice int
	initializers.DB.Table("book").Where("id IN (?)", orderInput.BookID).Select("SUM(total_price)").Scan(&totalPrice)
	order := models.Order{
		EmployeeID: orderInput.EmployeeID,
		OrderDate:  time.Now().Format("2006-01-02"),
		TotalPrice: totalPrice,
	}

	result := initializers.DB.Create(&order)
	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}
	// Return the order
	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

// ListOrders gets all the order
func ListOrders(c *gin.Context) {
	// Get all the orders
	var orders []models.Order
	var filter models.OrderFilter
	if err := c.BindQuery(&filter); err != nil {
		format_errors.InternalServerError(c)
		return
	}

	preloadFunc := func(query *gorm.DB) *gorm.DB {
		return query.Preload("Book", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, title, price, category, qty")
		}).Preload("Employee", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		})
	}

	result, err := pagination.Paginate(initializers.DB, filter.Page, filter.Limit, preloadFunc, &orders)
	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}

// GetOrder finds a order by ID
func GetOrder(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the order
	var order models.Order
	result := initializers.DB.Preload("Book", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title, price, category, qty")
	}).Preload("Employee", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).First(&order, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the order
	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

// Update Order
func UpdateOrder(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the data from request body
	var orderInput models.OrderRequest
	if err := c.ShouldBindJSON(&orderInput); err != nil {
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

	// Find the order by id
	var order models.Order
	result := initializers.DB.First(&order, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}
	//recalculate total price
	var totalPrice int
	initializers.DB.Table("book").Where("id IN (?)", orderInput.BookID).Select("SUM(total_price)").Scan(&totalPrice)

	// Prepare data to update
	updateOrder := models.Order{
		EmployeeID: orderInput.EmployeeID,
		OrderDate:  time.Now().Format("2006-01-02"),
		TotalPrice: totalPrice,
	}

	// Update the order
	finalResult := initializers.DB.Model(&order).Updates(&updateOrder)
	if finalResult.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the order
	c.JSON(http.StatusOK, gin.H{
		"order": updateOrder,
	})
}

// Soft Delete Order
func DeleteOrder(c *gin.Context) {
	// Get the id from the url
	id := c.Param("id")
	var order models.Order

	result := initializers.DB.First(&order, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the order
	initializers.DB.Delete(&order)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The order has been deleted successfully",
	})
}

// Permanent Delete Order
func PermanentlyDeleteOrder(c *gin.Context) {
	// Get id from url
	id := c.Param("id")
	var order models.Order

	// Find the order
	if err := initializers.DB.Unscoped().First(&order, id).Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the order
	initializers.DB.Unscoped().Delete(&order)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"message": "The order has been deleted permanently",
	})
}
