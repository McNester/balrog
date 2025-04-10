package main

import (
	"barlog/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	inventory_handler := handlers.NewInvenotryHandler()
	order_handler := handlers.NewOrderHandler()

	orders := router.Group("/orders")

	{
		orders.Any("/*", inventory_handler.Proxy)
	}

	products := router.Group("/productcs")

	{
		products.Any("/*", order_handler.Proxy)
	}

	// router.Any("/products/*", inventory_handler.Proxy)
	// router.Any("/orders/*", order_handler.Proxy)

	router.Run("0.0.0.0:8080")

}
