package models

import "github.com/shopspring/decimal"

// Currency @doc https://github.com/rubykube/peatio/blob/master/app/models/currency.rb
type Currency struct {
	BaseModel
	Symbol      string
	DepositFee  decimal.Decimal `json:"deposit_fee" sql:"DECIMAL(32,16)"`
	WithdrawFee decimal.Decimal `json:"withdraw_fee" sql:"DECIMAL(32,16)"`
}
