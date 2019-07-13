package dtos

import (
	"github.com/FlowerWrong/exchange/models"
	"github.com/shopspring/decimal"
)

// OrderBookDTO ...
type OrderBookDTO struct {
	// ID        uint      `json:"id"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	models.BaseModel
	UserID    uint            `json:"user_id"`
	Symbol    string          `json:"symbol"`
	FundID    uint            `json:"fund_id"`
	Status    uint            `json:"status"`
	OrderType string          `json:"order_type"`
	Side      string          `json:"side"`
	Volume    decimal.Decimal `json:"volume"`
	Price     decimal.Decimal `json:"price"`
}
