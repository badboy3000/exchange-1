package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/models"
	"github.com/FlowerWrong/exchange/services"
)

// QueueName ...
const QueueName = "trades_queue"

func main() {
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
					log.Println(orderBook)
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
