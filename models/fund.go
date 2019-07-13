package models

// Fund ...
type Fund struct {
	BaseModel
	Name            string
	Symbol          string   // eg btc_usd eth_usd
	RightCurrency   Currency `gorm:"foreignkey:RightCurrencyID"`
	RightCurrencyID uint64
	LeftCurrency    Currency `gorm:"foreignkey:LeftCurrencyID"`
	LeftCurrencyID  uint64
	LimitRate       float64 `json:"limit_rate" sql:"DECIMAL(32,16)"`
	MarketRate      float64 `json:"market_rate" sql:"DECIMAL(32,16)"`
}
