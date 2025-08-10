package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/n1x9s/L0/internal/db"
	"github.com/n1x9s/L0/internal/routers"
)

func main() {
	err := db.InitDB()
	if err != nil {
		return
	}

	r := gin.Default()
	r.Use(cors.Default())
	routers.RegisterOrderRoutes(r)

	r.Run()

}
