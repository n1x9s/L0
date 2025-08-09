package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/n1x9s/learnBasic/internal/handlers"
)

func RegisterOrderRoutes(r *gin.Engine) {
	orders := r.Group("/orders")
	orders.GET("/", handlers.GetOrders)
	orders.GET("/:order_uid", handlers.GetOrderById)
}
