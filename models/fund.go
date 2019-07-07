package models

import "github.com/jinzhu/gorm"

// Fund ...
type Fund struct {
	gorm.Model
	Name            string
	Symbol          string   // eg btc_usd eth_usd
	RightCurrency   Currency `gorm:"foreignkey:RightCurrencyID"`
	RightCurrencyID int64
	LeftCurrency    Currency `gorm:"foreignkey:LeftCurrencyID"`
	LeftCurrencyID  int64
	LimitRate       float64 `json:"limit_rate" sql:"DECIMAL(32,16)"`
	MarketRate      float64 `json:"market_rate" sql:"DECIMAL(32,16)"`
}
