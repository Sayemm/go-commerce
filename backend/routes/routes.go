package routes

import (
	"go-commerce/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/products", handlers.GetProducts)
	router.POST("/cart", handlers.AddToCart)
	router.POST("/order", handlers.PlaceOrder)
}
