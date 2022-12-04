package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	config "sr-server/config"
	service "sr-server/service"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.Origins,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.SetTrustedProxies(config.Origins)
	saleRoutes := router.Group("api/sale")
	{
		saleRoutes.GET("get-sales", service.GetSaleItems)

	}
	itemRoutes := router.Group("api/item")
	itemRoutes.Use()
	{
		itemRoutes.GET("all", service.GetItems)
		itemRoutes.POST("user", service.GetItemsByUser)
		itemRoutes.POST("create", service.CreateItem)
	}
	userRoutes := router.Group("api/user")
	{
		userRoutes.POST("create", service.CreateUser)
		userRoutes.POST("login", service.LogIn)
	}

	router.Run("localhost:8000")
}
