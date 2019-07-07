package models

import (
	"github.com/jinzhu/gorm"
)

// OrderBook ...
type OrderBook struct {
	gorm.Model
	UserID    int64
	User      User
	Symbol    string
	FundID    int64
	Fund      Fund
	OrderType string  // market or limit
	Side      string  // Sell Buy
	Volume    float64 `json:"volume" sql:"DECIMAL(32,16)"`
	Price     float64 `json:"price" sql:"DECIMAL(32,16)"`
}
