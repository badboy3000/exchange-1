package main

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/FlowerWrong/exchange/services/quotation"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	hub := quotation.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		quotation.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(":8100", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
