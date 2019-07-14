package main

import (
	"github.com/FlowerWrong/exchange/actions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/resty.v1"

	_ "./docs"
)

// @title Exchange server API
// @version 1.0
// @description Exchange server API 1.0 swagger doc
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	resty.SetDebug(false)
	resty.SetRESTMode()
	resty.SetHeader("Accept", "application/json")

	router := gin.Default()
	router.GET("/orders", actions.OrderIndex)
	router.GET("/order_books", actions.OrderBookIndex)
	router.POST("/order_books", actions.OrderBookCreate)
	router.PUT("/order_books/:id", actions.OrderBookUpdate)
	router.DELETE("/order_books/:id", actions.OrderBookCancel)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8080")
}
