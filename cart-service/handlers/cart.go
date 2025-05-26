package handlers

import (
	"net/http"

	"github.com/Sayemm/cart-service/db"
	"github.com/Sayemm/cart-service/models"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var item models.CartItem
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("INSERT INTO cart_items (product_id, quantity) VALUES (?, ?)", item.ProductID, item.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func GetCart(c *gin.Context) {
	var items []models.CartItem
	err := db.DB.Select(&items, "SELECT * FROM cart_items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func RemoveFromCart(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM cart_items WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item Removed from Cart"})
}
