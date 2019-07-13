package models

import "github.com/shopspring/decimal"

// Account ...
type Account struct {
	BaseModel
	UserID     uint64
	User       User
	CurrencyID uint64
	Currency   Currency
	Balance    decimal.Decimal `json:"balance" sql:"DECIMAL(32,16)"`
	Locked     decimal.Decimal `json:"locked" sql:"DECIMAL(32,16)"`
}
