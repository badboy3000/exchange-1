package main

import (
	"log"

	"github.com/FlowerWrong/exchange/actions"
	"github.com/FlowerWrong/exchange/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=kingyang dbname=exchange_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var currency models.Currency
	log.Println(db.First(&currency, 1))

	router := gin.Default()
	router.GET("/orders", actions.OrderIndex)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
