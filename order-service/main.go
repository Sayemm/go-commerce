package main

import (
	"github.com/Sayemm/order-service/db"
	"github.com/Sayemm/order-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	r.POST("/order", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)

	r.Run(":8002")
}
