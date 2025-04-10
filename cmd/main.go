package main

import (
	"barlog/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	registerService(router, "orders")
	registerService(router, "products")

	router.Run("0.0.0.0:8080")

}

func registerService(r *gin.Engine, service string) {
	proxyGroup := r.Group("/" + service)
	{
		proxyGroup.Any("/*path", handlers.NewProxy(service))
		proxyGroup.Any("", handlers.NewProxy(service))
	}
}
