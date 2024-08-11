package models

type Book struct {
	ID       int    `gorm:"primaryKey"`
	Title    string `gorm:"type:text" json:"title"`
	Price    int    `gorm:"type:integer;default:0" json:"price"`
	Category string `gorm:"type:text" json:"category"`
	Qty      int    `gorm:"type:integer;default:0" json:"qty"`
}

type BookRequest struct {
	Title    string `query:"title" json:"title"`
	Price    int    `query:"price" json:"price"`
	Category string `query:"category" json:"category"`
	Qty      int    `query:"qty" json:"qty"`
}

type BookFilter struct {
	BookIDS  []int  `query:"book_ids" json:"book_ids"`
	Title    string `query:"title" json:"title"`
	Price    int    `query:"price" json:"price"`
	Category string `query:"category" json:"category"`
	Qty      int    `query:"qty" json:"qty"`
	Page     int    `query:"page" json:"page"`
	Limit    int    `query:"limit" json:"limit"`
	Search   string `query:"search" json:"search"`
}
