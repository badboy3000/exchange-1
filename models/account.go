package models

import (
	"github.com/jinzhu/gorm"
)

// Account ...
type Account struct {
	gorm.Model
	UserID     int64
	User       User
	CurrencyID int64
	Currency   Currency
	Balance    float64 `json:"balance" sql:"DECIMAL(32,16)"`
	Locked     float64 `json:"locked" sql:"DECIMAL(32,16)"`
}
