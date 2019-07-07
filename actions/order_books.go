package actions

import (
	"net/http"

	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/models"
	"github.com/FlowerWrong/exchange/utils"
	"github.com/gin-gonic/gin"
)

// OrderBookIndex ...
//
// @Summary order book list
// @Description get your order book list
// @Accept json
// @Produce json
// @Param symbol query string true "eg BTC_USD"
// @Success 200 {array} models.OrderBook
// @Router /order_books?symbol={symbol} [get]
func OrderBookIndex(c *gin.Context) {
	symbol := c.DefaultQuery("symbol", "all")
	var orderBooks []models.OrderBook
	if symbol == "all" {
		db.ORM().Where("user_id = ?", 1).Find(&orderBooks)
	} else {
		db.ORM().Where("symbol = ? and user_id = ?", symbol, 1).Find(&orderBooks)
	}

	c.JSON(http.StatusOK, orderBooks)
}

// OrderBookCreate ...
//
// @Summary new an order book
// @Description create an order book
// @Accept json
// @Produce json
// @Param req body models.OrderBook true "order book model"
// @Success 200 {object} models.OrderBook
// @Router /order_books [post]
func OrderBookCreate(c *gin.Context) {
	var orderBook models.OrderBook
	if err := c.ShouldBindJSON(&orderBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderBook.UserID = 1 // FIXME

	// 相同品种，相同价格得单合并
	var orderBookExist models.OrderBook
	result := db.ORM().Where("symbol = ? and price = ?", orderBook.Symbol, orderBook.Price).First(&orderBookExist)
	if result.RecordNotFound() {
		var fund models.Fund
		db.ORM().Where("symbol = ?", orderBook.Symbol).First(&fund)
		orderBook.FundID = fund.ID
		db.ORM().Create(&orderBook)
	} else {
		orderBookExist.Volume = orderBookExist.Volume + orderBook.Volume
		db.ORM().Save(&orderBookExist)
		orderBook = orderBookExist
	}

	c.JSON(http.StatusOK, orderBook)
}

// OrderBookUpdate ...
//
// @Summary update an order book
// @Description update an exist order book
// @Accept json
// @Produce json
// @Param id path integer true "eg 1"
// @Param req body models.OrderBook true "order book model"
// @Success 200 {object} models.OrderBook
// @Router /order_books/{id} [put]
func OrderBookUpdate(c *gin.Context) {
	id := c.Param("id")
	var orderBookUpdate models.OrderBook
	if err := c.ShouldBindJSON(&orderBookUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var orderBook models.OrderBook
	db.ORM().Where("id = ? and volume > 0", id).First(&orderBook)
	orderBook.Volume = orderBookUpdate.Volume
	orderBook.Price = orderBookUpdate.Price
	db.ORM().Save(&orderBook)
	c.JSON(http.StatusOK, orderBook)
}

// OrderBookCancel ...
//
// @Summary cancel an order book
// @Description cancel an exist order book
// @Accept json
// @Produce json
// @Param id path integer true "eg 1"
// @Success 200 {object} utils.APIRes
// @Router /order_books/{id} [delete]
func OrderBookCancel(c *gin.Context) {
	id := c.Param("id")
	var orderBook models.OrderBook
	db.ORM().Where("id = ?", id).First(&orderBook).Delete(&orderBook)
	c.JSON(http.StatusOK, utils.APIRes{0, "ok"})
}
