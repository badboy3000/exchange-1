package forms

import (
	"testing"

	"github.com/FlowerWrong/exchange/models"
	"github.com/devfeel/mapper"
	"github.com/shopspring/decimal"
)

func TestOrderBookFormMapper(t *testing.T) {
	obf := &OrderBookForm{
		Symbol:    "BTC_USD",
		OrderType: "limit",
		Side:      "Buy",
		Volume:    decimal.NewFromFloat(10.00),
		Price:     decimal.NewFromFloat(100.00),
	}
	ob := &models.OrderBook{}
	mapper.AutoMapper(obf, ob)
	t.Log(obf)
	t.Log(ob)
	if obf.Symbol != ob.Symbol {
		t.Fatal("Wrong done id")
	}
}
