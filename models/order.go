package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductId uint `gorm:"foreignKey" json:"product_id"`
	Quantity  uint `json:"quantity"`
}
