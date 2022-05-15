package models

type Cart struct {
	CartId    uint64 `gorm:"primary_key;auto_increment" json:"cartid"`
	ProductId uint64 `gorm:"not null"`
	Quantity  int    `gorm:"not null"`
}
