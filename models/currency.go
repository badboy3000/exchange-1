package models

import (
	"github.com/jinzhu/gorm"
)

// Currency @doc https://github.com/rubykube/peatio/blob/master/app/models/currency.rb
type Currency struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);unique_index"`
	Symbol      string  `gorm:"type:varchar(100);unique_index"`
	DepositFee  float64 `json:"deposit_fee" sql:"DECIMAL(32,16)"`
	WithdrawFee float64 `json:"withdraw_fee" sql:"DECIMAL(32,16)"`
}
