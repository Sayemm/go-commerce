package handlers

import (
	"go-commerce/db"
	"go-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

	err := db.DB.Select(&products, "SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func AddToCart(c *gin.Context) {
	var item models.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added to cart", "item": item})
}

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order Placed", "order": order})
}
