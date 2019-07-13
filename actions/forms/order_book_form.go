package forms

import "github.com/shopspring/decimal"

// OrderBookForm ...
type OrderBookForm struct {
	Symbol    string          `json:"symbol"`
	OrderType string          `json:"order_type"` // market or limit
	Side      string          `json:"side"`       // sell or buy
	Volume    decimal.Decimal `json:"volume"`
	Price     decimal.Decimal `json:"price"`
}
