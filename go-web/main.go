package main

import (
	//"fmt"
	"net/http"
	//"os"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	Date      string  `json:"date"`
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

var lastID int = 3

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

func Save(c *gin.Context) {
	var req Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	lastID++
	req.ID = lastID
	products = append(products, req)
	c.JSON(200, req)
}

func main() {
	router := gin.Default()

	gopher := router.Group("/products")
	{
		gopher.GET("", GetAll)
		gopher.GET("/:id", GetId)
	}

	r := gin.Default()

	r.POST("/products", Save)

	router.Run()
}
