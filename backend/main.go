package main

import (
	"go-commerce/db"
	"go-commerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
