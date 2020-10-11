package models

import (
	"time"

	"gorm.io/gorm"
)

// Product struct!
type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImgAlt      string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	Name        string  `json:"productname"`
	Description string
}

// TableName Product Model !
func (Product) TableName() string {
	return "products"
}

// Customer struct!
type Customer struct {
	gorm.Model
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"colunm:email" json:"email"`
	Password  string  `gorm:"column:password" json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

// TableName Customer Model !
func (Customer) TableName() string {
	return "customers"
}

// Order struct!
type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id" json:"customer_id"`
	ProductID    int       `gorm:"column:product_id" json:"product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
