package models

import (
	"github.com/jinzhu/gorm"
)

// OrderBook ...
type OrderBook struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	User      User
	Symbol    string `json:"symbol"`
	FundID    uint   `json:"fund_id"`
	Fund      Fund
	OrderType string  `json:"order_type"` // market or limit
	Side      string  `json:"side"`       // Sell Buy
	Volume    float64 `json:"volume" sql:"DECIMAL(32,16)"`
	Price     float64 `json:"price" sql:"DECIMAL(32,16)"`
}
