package models

// Currency @doc https://github.com/rubykube/peatio/blob/master/app/models/currency.rb
type Currency struct {
	BaseModel
	Symbol      string
	DepositFee  float64 `json:"deposit_fee" sql:"DECIMAL(32,16)"`
	WithdrawFee float64 `json:"withdraw_fee" sql:"DECIMAL(32,16)"`
}
