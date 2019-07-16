package dtos

import "github.com/shopspring/decimal"

// KlineDTO is kline data for tradingview
type KlineDTO struct {
	Time   string          `json:"time"`
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume decimal.Decimal `json:"volume"`
}
