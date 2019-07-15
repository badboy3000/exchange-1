package actions

import "github.com/gin-gonic/gin"

// KlineIndex ...
//
// @Summary kline data
// @Description get kline data of fund
// @Accept json
// @Produce json
// @Param symbol query string "eg BTC_USD"
// @Success 200 {array} models.OrderBook
// @Router /kline?symbol={symbol} [get]
func KlineIndex(c *gin.Context) {
}
