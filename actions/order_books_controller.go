package actions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FlowerWrong/exchange/actions/forms"
	"github.com/FlowerWrong/exchange/db"
	"github.com/FlowerWrong/exchange/dtos"
	"github.com/FlowerWrong/exchange/models"
	"github.com/FlowerWrong/exchange/services"
	"github.com/FlowerWrong/exchange/utils"
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

// OrderBookIndex ...
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
func OrderBookCreate(c *gin.Context) {
	var orderBookForm forms.OrderBookForm
	if err := c.ShouldBindJSON(&orderBookForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderBook := &models.OrderBook{}
	mapper.AutoMapper(&orderBookForm, orderBook)
	orderBook.UserID = 1 // FIXME

	account := &models.Account{}
	fund := models.Fund{}
	db.ORM().First(&fund, orderBook.FundID)
	orderBook.FundID = fund.ID
	if orderBook.OrderType == "market" {
		orderBook.Price = models.CurrentPrice(orderBook.Symbol) // 现价 TOOD 如果库里没有订单怎么办?
	}
	if orderBook.Side == "buy" {
		// BTC_USD 为例，购买动作即用USD买BTC，锁定账户的USD
		turnover := orderBook.Volume.Mul(orderBook.Price) // 单价 * 数量
		models.FindAccountByUserIDAndCurrencyID(db.ORM(), account, orderBook.UserID, fund.RightCurrencyID)
		if account.Balance.Sub(turnover).Sign() < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrWithoutEnoughMoney})
			return
		}
	} else {
		models.FindAccountByUserIDAndCurrencyID(db.ORM(), account, orderBook.UserID, fund.LeftCurrencyID)
		if account.Balance.Sub(orderBook.Volume).Sign() < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": models.ErrWithoutEnoughMoney})
			return
		}
	}

	// 相同品种，相同价格得单不合并

	orderBook.CreatedAt = time.Now()
	orderBook.UpdatedAt = time.Now()

	id, err := utils.NextID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orderBook.ID = id

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

	db.RabbitmqChannel().Publish(
		"",                             // exchange
		"exchange.matching.work.queue", // routing key
		false,                          // mandatory
		false,                          // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})

	orderBookDTO := &dtos.OrderBookDTO{}
	mapper.AutoMapper(orderBook, orderBookDTO)
	c.JSON(http.StatusOK, orderBookDTO)
}

// OrderBookUpdate ...
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
func OrderBookCancel(c *gin.Context) {
	id := c.Param("id")
	// do not update here
	var orderBook models.OrderBook
	db.ORM().Where("id = ?", id).First(&orderBook).Delete(&orderBook)

	// 发送给queue

	c.JSON(http.StatusOK, utils.APIRes{Code: 0, Message: "ok"})
}
