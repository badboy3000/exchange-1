package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/models"
	"github.com/FlowerWrong/exchange/services"
	"github.com/FlowerWrong/exchange/services/matching"
)

// QueueName ...
const QueueName = "trades_queue"

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	matchEngine := matching.NewOrderBook()
	for {
		if db.Redis().Exists(QueueName).Val() > 0 {
			if db.Redis().LLen(QueueName).Val() > 0 {
				jsonStr, err := db.Redis().LPop(QueueName).Result()
				if err != nil {
					log.Println(err)
					continue
				}
				var event services.Event
				if err := json.Unmarshal([]byte(jsonStr), &event); err != nil {
					panic(err)
				}

				switch event.Name {
				case "create_order_book":
					var orderBook models.OrderBook
					err = json.Unmarshal(event.Data, &orderBook)
					if err != nil {
						panic(err)
					}

					log.Println("===========交易前==========")
					log.Println(matchEngine)
					// backup order book depth to redis
					obJSON, err := matchEngine.MarshalJSON()
					if err != nil {
						panic(err)
					}
					db.Redis().Set("matching_order_book", string(obJSON), 0)
					log.Println("=====================")
					side := matching.Str2Side(orderBook.Side)
					if orderBook.OrderType == "limit" {
						done, partial, partialQty, err := matchEngine.ProcessLimitOrder(side, orderBook.StrID(), orderBook.Volume, orderBook.Price)
						if err != nil {
							panic(err)
						}
						log.Println(done, partial, partialQty)
						models.Transaction(&orderBook, done)
					} else if orderBook.OrderType == "market" {
						done, partial, partialQty, left, err := matchEngine.ProcessMarketOrder(side, orderBook.Volume)
						if err != nil {
							panic(err)
						}
						log.Println(done, partial, partialQty, left)
						models.Transaction(&orderBook, done)
					}

					log.Println("=====================")
					log.Println(matchEngine)
					log.Println("===========交易后==========")
				case "update_order_book":
					var orderBook models.OrderBook
					err = json.Unmarshal(event.Data, &orderBook)
					if err != nil {
						panic(err)
					}
					log.Println(orderBook)
				case "cancel_order_book":
					// TODO
				default:
					fmt.Printf("Default")
				}
			} else {
				time.Sleep(time.Duration(100) * time.Microsecond)
			}
		} else {
			time.Sleep(time.Duration(2) * time.Second)
		}
	}
}
