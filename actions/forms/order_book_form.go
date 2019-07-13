package forms

// OrderBookForm ...
type OrderBookForm struct {
	Symbol    string  `json:"symbol"`
	OrderType string  `json:"order_type"` // market or limit
	Side      string  `json:"side"`       // sell or buy
	Volume    float64 `json:"volume"`
	Price     float64 `json:"price"`
}
