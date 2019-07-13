package models

// Order ...
type Order struct {
	BaseModel
	UserID       uint
	User         User
	FundID       uint
	Fund         Fund
	OrderBookID  uint
	OrderBook    OrderBook
	Symbol       string
	OrderType    string  // market or limit
	Side         string  // sell or buy
	Volume       float64 `json:"volume" sql:"DECIMAL(32,16)"`
	Price        float64 `json:"price" sql:"DECIMAL(32,16)"`
	AveragePrice float64 `json:"average_price" sql:"DECIMAL(32,16)"`
	AskFee       float64 `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee       float64 `json:"bid_fee" sql:"DECIMAL(32,16)"`
}
