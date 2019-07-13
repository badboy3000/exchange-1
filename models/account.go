package models

// Account ...
type Account struct {
	BaseModel
	UserID     uint
	User       User
	CurrencyID uint
	Currency   Currency
	Balance    float64 `json:"balance" sql:"DECIMAL(32,16)"`
	Locked     float64 `json:"locked" sql:"DECIMAL(32,16)"`
}
