package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/FlowerWrong/exchange/services/matching"
	"github.com/shopspring/decimal"
)

func addDepth(orderBook *matching.OrderBook, prefix string, quantity decimal.Decimal) {
	for i := 50; i < 100; i = i + 10 {
		orderBook.ProcessLimitOrder(matching.Buy, fmt.Sprintf("%sbuy-%d", prefix, i), quantity, decimal.New(int64(i), 0))
	}

	for i := 100; i < 150; i = i + 10 {
		orderBook.ProcessLimitOrder(matching.Sell, fmt.Sprintf("%ssell-%d", prefix, i), quantity, decimal.New(int64(i), 0))
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	orderBook := matching.NewOrderBook()
	addDepth(orderBook, "05-", decimal.New(10, 0))
	addDepth(orderBook, "10-", decimal.New(10, 0))
	addDepth(orderBook, "15-", decimal.New(10, 0))
	fmt.Println(orderBook)
}
