package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	EmployeeID int    `gorm:"foreignkey:EmployeeID" json:"employee_id"`
	OrderDate  string `gorm:"type:date" json:"order_date"`
	Books      []Book `json:"books" gorm:"many2many:order_books;"`
	TotalPrice int    `gorm:"type:integer;default:0" json:"total_price"`
}

type OrderRequest struct {
	EmployeeID int    `gorm:"foreignkey:EmployeeID" json:"employee_id"`
	OrderDate  string `gorm:"type:date" json:"order_date"`
	BookID     []int  `gorm:"not null" json:"book_id"`
}

type OrderRespomse struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	OrderDate  string `json:"order_date"`
	Books      []Book `json:"books"`
	TotalPrice int    `gorm:"type:integer;default:0" json:"total_price"`
}

type OrderFilter struct {
	BookIDS    []int  `query:"book_ids" json:"book_ids"`
	EmployeeID int    `query:"employee_id" json:"employee_id"`
	OrderDate  string `query:"order_date" json:"order_date"`
	BookID     []int  `query:"book_id" json:"book_id"`
	TotalPrice int    `query:"total_price" json:"total_price"`
	Page       int    `query:"page" json:"page"`
	Limit      int    `query:"limit" json:"limit"`
	Search     string `query:"search" json:"search"`
}
