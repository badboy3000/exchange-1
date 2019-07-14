package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

// Account ...
type Account struct {
	BaseModel
	UserID     uint64
	User       User
	CurrencyID uint64
	Currency   Currency
	Balance    decimal.Decimal `json:"balance" sql:"DECIMAL(32,16)"`
	Locked     decimal.Decimal `json:"locked" sql:"DECIMAL(32,16)"`
}

// Lock without db update
func (a *Account) Lock(money decimal.Decimal) {
	a.Locked = a.Locked.Add(money)
}

// UnLock without db update
func (a *Account) UnLock(money decimal.Decimal) {
	a.Locked = a.Locked.Sub(money)
	a.Balance = a.Balance.Sub(money)
}

// FindAccountByUserIDAndCurrencyID ...
func FindAccountByUserIDAndCurrencyID(tx *gorm.DB, account *Account, userID, currencyID uint64) {
	tx.Where("user_id = ? and currency_id = ?", userID, currencyID).First(account)
}

// Settlement 账户结算
func Settlement(order *Order, tx *gorm.DB) {
	fund := &Fund{}
	tx.First(fund, order.FundID)
	if order.Side == "buy" {
		accountRight := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountRight, order.UserID, fund.RightCurrencyID)
		accountRight.UnLock(order.Volume)
		tx.Save(accountRight)

		accountLeft := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountLeft, order.UserID, fund.LeftCurrencyID)
		accountLeft.Balance = accountLeft.Balance.Add(order.Volume.Mul(order.Price))
		tx.Save(accountLeft)
	} else {
		accountLeft := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountLeft, order.UserID, fund.LeftCurrencyID)
		accountLeft.UnLock(order.Volume)
		tx.Save(accountLeft)

		accountRight := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountRight, order.UserID, fund.RightCurrencyID)
		accountRight.Balance = accountRight.Balance.Add(order.Volume.Mul(order.Price))
		tx.Save(accountRight)
	}
}
