package models

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
	OrderType   string  // market or limit
	Side        string  // sell or buy
	Volume      float64 `json:"volume" sql:"DECIMAL(32,16)"`
	Price       float64 `json:"price" sql:"DECIMAL(32,16)"`
	AskFee      float64 `json:"ask_fee" sql:"DECIMAL(32,16)"`
	BidFee      float64 `json:"bid_fee" sql:"DECIMAL(32,16)"`
}
