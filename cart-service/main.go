package main

import (
	"github.com/Sayemm/cart-service/db"
	"github.com/Sayemm/cart-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	r.POST("/cart", handlers.AddToCart)
	r.GET("/cart", handlers.GetCart)
	r.DELETE("/cart/:id", handlers.RemoveFromCart)

	r.Run(":8001")
}
