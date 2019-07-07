package models

import "github.com/jinzhu/gorm"

// Deposit ...
type Deposit struct {
	gorm.Model
	AccountID  int64
	Account    Account
	CurrencyID int64
	Currency   Currency
	Amount     float64 `json:"amount" sql:"DECIMAL(32,16)"`
	Fee        float64 `json:"fee" sql:"DECIMAL(32,16)"`
}
