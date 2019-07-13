package models

import "github.com/shopspring/decimal"

// OrderBook ...
type OrderBook struct {
	BaseModel
	UserID    uint64 `json:"user_id"`
	User      User
	Symbol    string `json:"symbol"`
	FundID    uint64 `json:"fund_id"`
	Fund      Fund
	Status    uint            `json:"status"`     // pending done cancel reject
	OrderType string          `json:"order_type"` // market or limit
	Side      string          `json:"side"`       // sell or buy
	Volume    decimal.Decimal `json:"volume" sql:"DECIMAL(32,16)"`
	Price     decimal.Decimal `json:"price" sql:"DECIMAL(32,16)"`
}
