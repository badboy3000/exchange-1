package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/FlowerWrong/exchange/services/kline"
)

// @doc https://github.com/beimingio/peatio/blob/master/lib/daemons/k.rb
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println(kline.K1("btc_usdt", time.Now().Add(-10*time.Minute)))

	// for {
	// 	var funds []models.Fund
	// 	db.ORM().Find(&funds)

	// 	for _, fund := range funds {
	// 		log.Println(fund)
	// 	}

	// 	time.Sleep(time.Duration(10) * time.Second)
	// }
}
