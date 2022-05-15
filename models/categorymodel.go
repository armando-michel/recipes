package models

import (
	"time"
)

type Category struct {
	CategoryId    uint64    `gorm:"primary_key;auto_increment" json:"categoryid"`
	Name          string    `gorm:"size:100;not null;" json:"name"`
	SubcategoryId uint64    `gorm:"not null" json:"subcategory"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
