package models

import "time"

type Order struct {
	OrderId     uint64    `gorm:"primary_key;auto_increment" json:"orderid"`
	ClientId    uint64    `gorm:"not null"`
	OrderDate   time.Time `gorm:"not null"`
	ShipDate    time.Time `gorm:"not null"`
	OrderStatus string    `gorm:"not null"`
}

type OrderDetails struct {
	OrderId     uint64  `gorm:"not null"`
	ProductId   uint64  `gorm:"not null"`
	OrderNumber int     `gorm:"not null"`
	ProductName string  `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Discount    int     `gorm:"not"`
	Total       int
}
