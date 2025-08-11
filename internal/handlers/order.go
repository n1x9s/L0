package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n1x9s/L0/internal/cache"
)

func GetOrders(c *gin.Context) {
	orders := cache.Cache.GetAllOrders()
	c.JSON(http.StatusOK, orders)
}

func GetOrderById(c *gin.Context) {
	orderUID := c.Param("order_uid")

	// Сначала ищем в кеше
	order, exists := cache.Cache.GetOrder(orderUID)
	if exists {
		c.JSON(http.StatusOK, order)
		return
	}

	// Если в кеше нет, пытаемся загрузить из БД и добавить в кеш
	order, err := cache.Cache.GetOrderFromDBAndCache(orderUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
