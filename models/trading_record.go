package models

import "github.com/shopspring/decimal"

// TradingRecord ...
type TradingRecord struct {
	BaseModel
	UserID      uint64
	User        User
	FundID      uint64
	Fund        Fund
	Symbol      string
	OrderBookID uint64
	OrderBook   OrderBook
	OrderID     uint64
	Order       Order
	OrderType   string          // market or limit
	Side        string          // sell or buy
	Volume      decimal.Decimal `json:"volume" sql:"DECIMAL(32,16)"`
	Price       decimal.Decimal `json:"price" sql:"DECIMAL(32,16)"`
	AskFee      decimal.Decimal `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee      decimal.Decimal `json:"bid_fee" sql:"DECIMAL(32,16)"`
}
