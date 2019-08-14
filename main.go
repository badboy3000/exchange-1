package main

import (
	"github.com/FlowerWrong/exchange/actions"
	"github.com/FlowerWrong/exchange/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/resty.v1"
)

func main() {
	resty.SetDebug(false)
	resty.SetRESTMode()
	resty.SetHeader("Accept", "application/json")

	router := gin.Default()
	router.GET("/orders", actions.OrderIndex)
	router.GET("/order_books", actions.OrderBookIndex)
	router.POST("/order_books", actions.OrderBookCreate)
	router.PUT("/order_books/:id", actions.OrderBookUpdate)
	router.DELETE("/order_books/:id", actions.OrderBookCancel)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// rabbitmq
	db.RabbitmqChannel()
	db.DeclareMatchingWorkQueue()

	router.Run(":8080")
}
