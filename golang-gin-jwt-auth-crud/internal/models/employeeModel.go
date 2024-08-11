package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID            int    `gorm:"primaryKey"`
	Name          string `gorm:"unique;not null" json:"name"`
	Email         string `gorm:"unique;not null" json:"email"`
	Address       string `gorm:"type:varchar(255)" json:"address"`
	Status        string `json:"status" gorm:"type:varchar(255)"`
	Handphone     string `gorm:"type:varchar(14)" json:"handphone"`
	Gender        string `gorm:"type:varchar(255)" json:"gender"`
	BirthPlace    string `gorm:"type:varchar(255)" json:"birth_place"`
	BirthDate     string `gorm:"type:date;not null" json:"birth_date"`
	MaritalStatus string `gorm:"type:varchar(255)" json:"marital_status"`
}

type EmployeeRequest struct {
	Name          string `query:"name" json:"name"`
	Email         string `query:"email" json:"email"`
	Address       string `query:"address" json:"address"`
	Status        string `query:"status" json:"status"`
	Handphone     string `query:"handphone" json:"handphone"`
	Gender        string `query:"gender" json:"gender"`
	BirthPlace    string `query:"birth_place" json:"birth_place"`
	BirthDate     string `query:"birth_date" json:"birth_date"`
	MaritalStatus string `query:"marital_status" json:"marital_status"`
}

type EmployeeFilter struct {
	EmployeeIDS   []int  `query:"employee_ids" json:"employee_ids"`
	Name          string `query:"name" json:"name"`
	Email         string `query:"email" json:"email"`
	Address       string `query:"address" json:"address"`
	Status        string `query:"status" json:"status"`
	Handphone     string `query:"handphone" json:"handphone"`
	Gender        string `query:"gender" json:"gender"`
	BirthPlace    string `query:"birth_place" json:"birth_place"`
	BirthDate     string `query:"birth_date" json:"birth_date"`
	MaritalStatus string `query:"marital_status" json:"marital_status"`
	Page          int    `query:"page" json:"page"`
	Limit         int    `query:"limit" json:"limit"`
	Search        string `query:"search" json:"search"`
}
