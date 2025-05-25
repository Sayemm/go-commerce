package main

import (
	"github.com/Sayemm/product-service/db"
	"github.com/Sayemm/product-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/products", handlers.GetProducts)
	r.POST("products", handlers.AddProduct)

	r.Run(":8000")
}
