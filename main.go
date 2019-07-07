package main

import (
	"github.com/FlowerWrong/exchange/actions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/orders", actions.OrderIndex)
	router.GET("/order_books", actions.OrderBookIndex)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
