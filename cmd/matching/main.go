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
	addDepth(orderBook, "", decimal.New(2, 0))
	fmt.Println(orderBook)
	log.Println(orderBook.ProcessLimitOrder(matching.Buy, "order-b100", decimal.New(1, 0), decimal.New(100, 0)))

	// done, partial, partialQty, _ := orderBook.ProcessLimitOrder(matching.Buy, fmt.Sprintf("%sbuy-%d", "80", 80), decimal.New(1, 0), decimal.New(80, 0))
	// log.Println(done, partial, partialQty)

	// done, partial, partialQty, _ = orderBook.ProcessLimitOrder(matching.Buy, fmt.Sprintf("%sbuy-%d", "90", 90), decimal.New(5, 0), decimal.New(90, 0))
	// log.Println(done, partial, partialQty)

	// done, partial, partialQty, _ = orderBook.ProcessLimitOrder(matching.Sell, fmt.Sprintf("%sbuy-%d", "100", 100), decimal.New(1, 0), decimal.New(100, 0))
	// log.Println(done, partial, partialQty)

	// done, partial, partialQty, _ = orderBook.ProcessLimitOrder(matching.Sell, fmt.Sprintf("%sbuy-%d", "110", 110), decimal.New(5, 0), decimal.New(110, 0))
	// log.Println(done, partial, partialQty)

	// log.Println(orderBook.ProcessMarketOrder(matching.Buy, decimal.New(10, 0)))
	fmt.Println(orderBook)

	// done, partial, partialQty, _ = orderBook.ProcessLimitOrder(matching.Buy, "uinqueID", decimal.New(3, 0), decimal.New(120, 0))
	// log.Println(done, partial, partialQty)
	// fmt.Println(orderBook)
}
