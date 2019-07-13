package models

// Withdraw ...
type Withdraw struct {
	BaseModel
	AccountID  uint64
	Account    Account
	CurrencyID uint64
	Currency   Currency
	Amount     float64 `json:"amount" sql:"DECIMAL(32,16)"`
	Fee        float64 `json:"fee" sql:"DECIMAL(32,16)"`
}
