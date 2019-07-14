package models

import (
	"github.com/shopspring/decimal"
)

// Fund ...
type Fund struct {
	BaseModel
	Name            string
	Symbol          string   // eg btc_usd eth_usd
	RightCurrency   Currency `gorm:"foreignkey:RightCurrencyID"`
	RightCurrencyID uint64
	LeftCurrency    Currency `gorm:"foreignkey:LeftCurrencyID"`
	LeftCurrencyID  uint64
	LimitRate       decimal.Decimal `json:"limit_rate" sql:"DECIMAL(32,16)"`
	MarketRate      decimal.Decimal `json:"market_rate" sql:"DECIMAL(32,16)"`
}
