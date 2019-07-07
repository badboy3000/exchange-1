package models

import "github.com/jinzhu/gorm"

// Order ...
type Order struct {
	gorm.Model
	UserID    uint
	User      User
	FundID    uint
	Fund      Fund
	Symbol    string
	OrderType string  // market or limit
	Side      string  // Sell Buy
	Volume    float64 `json:"volume" sql:"DECIMAL(32,16)"`
	Price     float64 `json:"price" sql:"DECIMAL(32,16)"`
	AskFee    float64 `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee    float64 `json:"bid_fee" sql:"DECIMAL(32,16)"`
}
