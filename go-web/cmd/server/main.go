package main

import (
	//"fmt"
	"net/http"
	//"os"
	"github.com/gin-gonic/gin"
	"strconv"
)





func Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "aaaaa" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req Product

		if err := c.ShouldBindJSON(&req); err != nil {
			if req.Name == "" || req.Color == "" || req.Price == 0 || req.Stock == 0 || req.Code == "" || req.Date == "" {
				if req.Name == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "¡El campo Nombre es requerido!"})
				}
				if req.Color == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "¡El campo Color es requerido!"})
				}
				if req.Price == 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "¡El campo Precio es requerido!"})
				}
				if req.Stock == 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "¡El campo Stock es requerido!"})
				}
				if req.Code == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "¡El campo Codigo es requerido!"})
				}
				if req.Date == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "¡El campo Fecha es requerido!"})
				} else {
					c.JSON(404, gin.H{
						"error": err.Error(),
					})
					return
				}
			}
		}
		//lastID++
		req.ID = len(products) + 1
		products = append(products, req)
		c.JSON(http.StatusOK, req)
	}
}

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
