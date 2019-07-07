package models

import "github.com/jinzhu/gorm"

// Deposit ...
type Deposit struct {
	gorm.Model
	UserID     int64
	User       User
	CurrencyID int64
	Currency   Currency
	Amount     float64 `json:"amount" sql:"DECIMAL(32,16)"`
	Fee        float64 `json:"fee" sql:"DECIMAL(32,16)"`
}
