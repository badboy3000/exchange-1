package models

import (
	"github.com/jinzhu/gorm"
)

// OrderBook ...
type OrderBook struct {
	gorm.Model
	Symbol string
	Side   string  // Sell Buy
	Size   float64 `json:"size" sql:"DECIMAL(32,16)"`
	Price  float64 `json:"price" sql:"DECIMAL(32,16)"`
}
