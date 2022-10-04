package main

import (
	//"fmt"
	"net/http"
	//"os"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Product struct {
	ID        int     `json:"id" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	Color     string  `json:"color" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	Code      string  `json:"code" binding:"required"`
	Published bool    `json:"published" binding:"required"`
	Date      string  `json:"date" binding:"required"`
}

var products = []Product{
	{
		ID:        1,
		Name:      "lata",
		Color:     "blanco",
		Price:     20.00,
		Stock:     300,
		Code:      "a00a",
		Published: false,
		Date:      "20-09-2022",
	},
	{
		ID:        2,
		Name:      "bolsa",
		Color:     "gris",
		Price:     3.00,
		Stock:     20,
		Code:      "b00a",
		Published: true,
		Date:      "20-09-2022",
	},
	{
		ID:        3,
		Name:      "cajón",
		Color:     "marrón",
		Price:     40.00,
		Stock:     90,
		Code:      "c04j",
		Published: false,
		Date:      "20-09-2022",
	},
}

//var lastID int

func GetAll(c *gin.Context) {
	var filtered []Product
	query := c.Query("name")
	for _, p := range products {
		if query == p.Name {
			filtered = append(filtered, p)
		}
	}

	c.JSON(http.StatusOK, filtered)
}

func GetId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(404, "Información del producto ¡No existe! \n")
	} else {
		c.String(200, "Información del producto %s, nombre: %s \n", c.Param("id"), id)
	}
}

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
