package models

import (
	"time"
)

type Product struct {
	ProductId   uint64    `gorm:"primary_key;auto_increment" json:"productid"`
	Name        string    `gorm:"size:100;not null;" json:"name"`
	Description string    `gorm:"size:100" json:"description"`
	Price       float64   `gorm:"not null" json:"sub"`
	CategoryId  uint64    `gorm:"not null" json:"categoryid"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ProductVariation struct {
	VariationId   uint64
	ProductId     uint64
	VariationName string
}

type ProductVariationOption struct {
	Id          uint64
	VariationId uint64
	Name        string
}

//  Product variation																Product Variation Options
// id 			variation name  product Id							id  		variationid       name
// 1				size 						2                       1       1									S
// 2  			color           3  			                2       1     						M
// 																									3       1                 L

// 																									4       2                 Blue
// 																									5       2                 Pink
// 																									5       2                 Green
