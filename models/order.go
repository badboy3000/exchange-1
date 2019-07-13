package models

import "github.com/shopspring/decimal"

// Order ...
type Order struct {
	BaseModel
	UserID       uint64
	User         User
	FundID       uint64
	Fund         Fund
	Symbol       string
	OrderType    string          // market or limit
	Side         string          // sell or buy
	Volume       decimal.Decimal `json:"volume" sql:"DECIMAL(32,16)"`
	Price        decimal.Decimal `json:"price" sql:"DECIMAL(32,16)"`
	AveragePrice decimal.Decimal `json:"average_price" sql:"DECIMAL(32,16)"`
	AskFee       decimal.Decimal `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee       decimal.Decimal `json:"bid_fee" sql:"DECIMAL(32,16)"`
}
