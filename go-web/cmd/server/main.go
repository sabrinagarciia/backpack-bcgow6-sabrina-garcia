package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrinagarciia/backpack-bcgow6-sabrina-garcia/go-web/internal/products"
	"github.com/sabrinagarciia/backpack-bcgow6-sabrina-garcia/go-web/cmd/server/handler"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	router := gin.Default()

	gopher := router.Group("/products")
	{
		gopher.GET("", p.GetAll)
		//gopher.GET("/:id", p.GetId)
	}

	router.POST("products", p.Save())
	router.POST("products", p.Update())
	
	router.Run()
}
