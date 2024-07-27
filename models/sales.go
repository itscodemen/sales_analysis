package models

import (
	"time"
)

type Order struct {
	ID              uint   `gorm:"primaryKey"`
	OrderID         string `gorm:"unique"`
	ProductID       string
	CustomerID      string
	ProductName     string
	Category        string
	Region          string
	DateOfSale      time.Time
	QuantitySold    int
	UnitPrice       float64
	Discount        float64
	ShippingCost    float64
	PaymentMethod   string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	ProductID   string `gorm:"unique"`
	ProductName string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Customer struct {
	ID              uint   `gorm:"primaryKey"`
	CustomerID      string `gorm:"unique"`
	CustomerName    string
	CustomerEmail   string
	CustomerAddress string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
