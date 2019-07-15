package models

import (
	"github.com/FlowerWrong/exchange/db"
	"github.com/shopspring/decimal"
)

// Order ...
type Order struct {
	BaseModel
	UserID               uint64
	User                 User
	FundID               uint64
	Fund                 Fund
	OrderBookID          uint64
	OrderBook            OrderBook
	OtherSideOrderBookID uint64
	OtherSideOrderID     uint64
	Symbol               string
	OrderType            string          // market or limit
	Side                 string          // sell or buy
	Volume               decimal.Decimal `json:"volume" sql:"DECIMAL(32,16)"`
	Price                decimal.Decimal `json:"price" sql:"DECIMAL(32,16)"`
	AveragePrice         decimal.Decimal `json:"average_price" sql:"DECIMAL(32,16)"`
	AskFee               decimal.Decimal `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee               decimal.Decimal `json:"bid_fee" sql:"DECIMAL(32,16)"`
}

// CurrentPrice 返回最新成交价
func CurrentPrice(symbol string) decimal.Decimal {
	var order Order
	db.ORM().Where("symbol = ?", symbol).Last(&order)
	return order.Price
}
