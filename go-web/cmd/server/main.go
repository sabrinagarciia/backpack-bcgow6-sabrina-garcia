package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	gopher := router.Group("/products")
	{
		gopher.GET("", GetAll)
		gopher.GET("/:id", GetId)
	}

	router.POST("products", Save())

	router.Run()
}
