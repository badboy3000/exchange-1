package models

import "errors"

var (
	ErrWithoutEnoughMoney = errors.New("orderbook: without enough money")
)
