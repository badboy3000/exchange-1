package models

import (
	"github.com/FlowerWrong/exchange/db"
	"github.com/shopspring/decimal"
)

// Order ...
type Order struct {
	BaseModel
	BidUserID      uint64
	BidUser        User
	AskUserID      uint64
	AskUser        User
	FundID         uint64
	Fund           Fund
	BidOrderBookID uint64
	BidOrderBook   OrderBook
	AskOrderBookID uint64
	AskOrderBook   OrderBook
	Symbol         string
	Volume         decimal.Decimal `json:"volume" sql:"DECIMAL(32,16)"`
	Price          decimal.Decimal `json:"price" sql:"DECIMAL(32,16)"`
	AskFee         decimal.Decimal `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee         decimal.Decimal `json:"bid_fee" sql:"DECIMAL(32,16)"`
}

// CurrentPrice 返回最新成交价
func CurrentPrice(symbol string) decimal.Decimal {
	var order Order
	db.ORM().Where("symbol = ?", symbol).Last(&order)
	return order.Price
}
