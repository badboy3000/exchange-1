package actions

import (
	"log"
	"net/http"

	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/models"
	"github.com/gin-gonic/gin"
)

// OrderBookIndex ...
func OrderBookIndex(c *gin.Context) {
	var currency models.Currency
	log.Println(db.ORM().First(&currency, 1))

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// OrderBookCreate ...
func OrderBookCreate(c *gin.Context) {

}

// OrderBookUpdate ...
func OrderBookUpdate(c *gin.Context) {

}

// OrderBookCancel ...
func OrderBookCancel(c *gin.Context) {

}
