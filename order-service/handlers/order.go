package handlers

import (
	"net/http"
	"time"

	"github.com/Sayemm/order-service/db"
	"github.com/Sayemm/order-service/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.CreatedAt = time.Now()

	_, err := db.DB.Exec(
		"INSERT INTO orders (product_id, quantity, created_at) VALUES (?, ?, ?)",
		order.ProductID, order.Quantity, order.CreatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order Place successfully"})
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	err := db.DB.Select(&orders, "SELECT * FROM orders ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
