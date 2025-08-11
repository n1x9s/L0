package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/n1x9s/L0/internal/cache"
	"github.com/n1x9s/L0/internal/db"
	"github.com/n1x9s/L0/internal/routers"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
		return
	}

	err = cache.InitCache()
	if err != nil {
		log.Fatalf("Ошибка инициализации кеша: %v", err)
		return
	}

	r := gin.Default()
	r.Use(cors.Default())
	routers.RegisterOrderRoutes(r)

	err = r.Run()
	if err != nil {
		return
	}
}
