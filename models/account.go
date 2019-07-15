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
	turnover := order.Volume.Mul(order.Price)
	// BTC_USD 为例，购买动作即用USD买BTC
	{
		// USD减少
		accountRight := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountRight, order.BidUserID, fund.RightCurrencyID)
		accountRight.UnLock(turnover)
		tx.Save(accountRight)

		// BTC增加
		accountLeft := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountLeft, order.BidUserID, fund.LeftCurrencyID)
		accountLeft.Balance = accountLeft.Balance.Add(order.Volume)
		tx.Save(accountLeft)
	}
	{
		// USD增加
		accountLeft := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountLeft, order.AskUserID, fund.LeftCurrencyID)
		accountLeft.UnLock(turnover)
		tx.Save(accountLeft)

		// BTC减少
		accountRight := &Account{}
		FindAccountByUserIDAndCurrencyID(tx, accountRight, order.AskUserID, fund.RightCurrencyID)
		accountRight.Balance = accountRight.Balance.Add(order.Volume)
		tx.Save(accountRight)
	}
}
