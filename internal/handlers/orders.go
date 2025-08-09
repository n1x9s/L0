package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n1x9s/learnBasic/internal/db"
	"github.com/n1x9s/learnBasic/internal/models"
)

func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := db.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orders not found"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderById(c *gin.Context) {
	orderUID := c.Param("order_uid")

	var order models.Order
	if err := db.DB.Preload("Items").First(&order, "order_uid = ?", orderUID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
