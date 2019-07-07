package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AccountIndex ...
func AccountIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// AccountShow ...
func AccountShow(c *gin.Context) {
	currency := c.Param("currency")
	c.JSON(http.StatusOK, gin.H{
		"status": currency,
	})
}
