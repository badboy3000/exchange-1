package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)

// @doc https://chai2010.cn/advanced-go-programming-book/ch6-cloud/ch6-01-dist-id.html
var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

// curl 127.0.0.1:8090
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		id, err := sf.NextID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, sonyflake.Decompose(id))
	})

	router.Run(":8090")
}
