package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderIndex ...
func OrderIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
