package models

// OrderBook ...
type OrderBook struct {
	BaseModel
	UserID    uint `json:"user_id"`
	User      User
	Symbol    string `json:"symbol"`
	FundID    uint   `json:"fund_id"`
	Fund      Fund
	Status    uint    `json:"status"`     // pending done cancel reject
	OrderType string  `json:"order_type"` // market or limit
	Side      string  `json:"side"`       // sell or buy
	Volume    float64 `json:"volume" sql:"DECIMAL(32,16)"`
	Price     float64 `json:"price" sql:"DECIMAL(32,16)"`
}
