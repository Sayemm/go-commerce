package handlers

import (
	"net/http"

	"github.com/Sayemm/product-service/db"
	"github.com/Sayemm/product-service/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	err := db.DB.Select(&products, "SELECT * FROM products")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var p models.Product
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)", p.Name, p.Price, p.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product added"})
}
