package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Me ...
func Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
