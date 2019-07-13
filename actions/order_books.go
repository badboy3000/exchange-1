package actions

import (
	"encoding/json"
	"net/http"

	"github.com/FlowerWrong/exchange/dtos"

	"github.com/FlowerWrong/exchange/actions/forms"
	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/models"
	"github.com/FlowerWrong/exchange/services"
	"github.com/FlowerWrong/exchange/utils"
	"github.com/devfeel/mapper"
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
// @Param req body forms.OrderBookForm true "order book form"
// @Success 200 {object} dtos.OrderBookDTO
// @Router /order_books [post]
func OrderBookCreate(c *gin.Context) {
	var orderBookForm forms.OrderBookForm
	if err := c.ShouldBindJSON(&orderBookForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderBook := &models.OrderBook{}
	mapper.AutoMapper(&orderBookForm, orderBook)
	orderBook.UserID = 1 // FIXME

	// 相同品种，相同价格得单不合并

	var fund models.Fund
	db.ORM().Where("symbol = ?", orderBook.Symbol).First(&fund)
	orderBook.FundID = fund.ID
	db.ORM().Create(&orderBook)

	// 发送给queue
	b, err := json.Marshal(orderBook)
	if err != nil {
		panic(err)
	}
	raw := json.RawMessage(b)
	event := &services.Event{Name: "create_order_book", Data: raw}
	data, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	db.Redis().RPush("trades_queue", string(data))

	orderBookDTO := &dtos.OrderBookDTO{}
	mapper.AutoMapper(orderBook, orderBookDTO)
	c.JSON(http.StatusOK, orderBookDTO)
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

	// do not update here
	var orderBook models.OrderBook
	db.ORM().Where("id = ? and volume > 0", id).First(&orderBook)
	orderBook.Volume = orderBookUpdate.Volume
	orderBook.Price = orderBookUpdate.Price
	db.ORM().Save(&orderBook)

	// 发送给queue

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
	// do not update here
	var orderBook models.OrderBook
	db.ORM().Where("id = ?", id).First(&orderBook).Delete(&orderBook)

	// 发送给queue

	c.JSON(http.StatusOK, utils.APIRes{Code: 0, Message: "ok"})
}
