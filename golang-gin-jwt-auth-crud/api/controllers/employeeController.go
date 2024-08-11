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

// GetAll Employee
func ListEmployee(c *gin.Context) {
	var allEmployee []models.Employee

	var filter models.EmployeeFilter
	if err := c.BindQuery(&filter); err != nil {
		format_errors.InternalServerError(c)
		return
	}

	result, err := pagination.Paginate(initializers.DB, filter.Page, filter.Limit, nil, &allEmployee)

	if err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the books
	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}

// Create Employee
func CreateEmployee(c *gin.Context) {
	// Get data from request
	var employeeInput models.EmployeeRequest
	err := c.ShouldBindJSON(&employeeInput)

	// Validate the data
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errors),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check email
	if validations.IsExistValue("employees", "email", employeeInput.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"validations": map[string]interface{}{
				"Email": "The email is already used",
			},
		})
		return
	}

	//check phone number
	if validations.IsExistValue("employees", "handphone", employeeInput.Handphone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"validations": map[string]interface{}{
				"Handphone": "The phone number is already used",
			},
		})
		return
	}

	employee := models.Employee{
		Name:          employeeInput.Name,
		Email:         employeeInput.Email,
		Address:       employeeInput.Address,
		Handphone:     employeeInput.Handphone,
		Gender:        employeeInput.Gender,
		BirthPlace:    employeeInput.BirthPlace,
		BirthDate:     employeeInput.BirthDate,
		MaritalStatus: employeeInput.MaritalStatus,
	}

	result := initializers.DB.Create(&employee)
	if result.Error != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the employee
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

// Get Employee By Id
func GetEmployee(c *gin.Context) {
	// Get the employee id from url
	id := c.Param("id")

	// Find the employee
	var employee models.Employee
	result := initializers.DB.First(&employee, id)

	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the employee
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

// Update Employee
func UpdateEmployee(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	var employeeInput models.EmployeeRequest
	// Validate in request
	err := c.ShouldBindJSON(&employeeInput)
	// Validate the data
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"validations": validations.FormatValidationErrors(errors),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the employee
	var employee models.Employee
	result := initializers.DB.First(&employee, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Check unique validation
	if validations.IsUniqueValue("employees", "email", employeeInput.Email, employee.ID) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Email": "The email is already used!",
			},
		})

		return
	}

	if validations.IsUniqueValue("employees", "handphone", employeeInput.Handphone, employee.ID) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Handphone": "The phone number is already used!",
			},
		})

		return
	}

	updateEmployee := models.Employee{
		Name:          employeeInput.Name,
		Email:         employeeInput.Email,
		Address:       employeeInput.Address,
		Handphone:     employeeInput.Handphone,
		Gender:        employeeInput.Gender,
		BirthPlace:    employeeInput.BirthPlace,
		BirthDate:     employeeInput.BirthDate,
		MaritalStatus: employeeInput.MaritalStatus,
	}

	// Update the book record
	result = initializers.DB.Model(&employee).Updates(updateEmployee)
	if err := result.Error; err != nil {
		format_errors.InternalServerError(c)
		return
	}

	// Return the employee
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
}

// Soft deletes a employee by id
func DeleteEmployee(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")
	var employee models.Employee

	// Find the employee
	result := initializers.DB.First(&employee, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Delete the employee
	initializers.DB.Delete(&employee)

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The employee has been deleted successfully",
	})
}

// Permanent deletes a employee by id
func DeleteEmployeePermanent(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")

	// Delete the employee
	result := initializers.DB.Unscoped().Delete(&models.Employee{}, id)
	if err := result.Error; err != nil {
		format_errors.RecordNotFound(c, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The employee has been deleted permanently",
	})
}
