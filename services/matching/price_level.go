package matching

import "github.com/shopspring/decimal"

// PriceLevel contains price and volume in depth
type PriceLevel struct {
	Price    decimal.Decimal `json:"price"`
	Quantity decimal.Decimal `json:"quantity"`
}
