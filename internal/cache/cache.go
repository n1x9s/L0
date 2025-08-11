package cache

import (
	"log"
	"sync"

	"github.com/n1x9s/L0/internal/db"
	"github.com/n1x9s/L0/internal/models"
)

type OrderCache struct {
	cache map[string]*models.Order
	mutex sync.RWMutex
}

var Cache *OrderCache

func init() {
	Cache = &OrderCache{
		cache: make(map[string]*models.Order),
	}
}

func InitCache() error {
	Cache.mutex.Lock()
	defer Cache.mutex.Unlock()

	var orders []models.Order
	if err := db.DB.Preload("Items").Find(&orders).Error; err != nil {
		log.Printf("Ошибка загрузки данных из БД для кеша: %v", err)
		return err
	}

	for i := range orders {
		Cache.cache[orders[i].OrderUID] = &orders[i]
	}

	log.Printf("Кеш инициализирован, загружено %d заказов", len(orders))
	return nil
}

func (c *OrderCache) GetOrder(orderUID string) (*models.Order, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	order, exists := c.cache[orderUID]
	return order, exists
}

func (c *OrderCache) SetOrder(order *models.Order) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[order.OrderUID] = order
}

func (c *OrderCache) GetAllOrders() []*models.Order {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	orders := make([]*models.Order, 0, len(c.cache))
	for _, order := range c.cache {
		orders = append(orders, order)
	}

	return orders
}

func (c *OrderCache) GetOrderFromDBAndCache(orderUID string) (*models.Order, error) {
	var order models.Order
	if err := db.DB.Preload("Items").First(&order, "order_uid = ?", orderUID).Error; err != nil {
		return nil, err
	}

	c.SetOrder(&order)
	return &order, nil
}
