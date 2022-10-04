package main

import (
	//"fmt"
	"net/http"
	//"os"
	"github.com/gin-gonic/gin"
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

func GetAll(c *gin.Context) {
	//c.JSON(http.StatusOK, products)
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
	//prod, err := products[c.Param("id")]

	// if err {
	// 	c.String(200, "Información del producto %s, nombre: %s \n", c.Param("id"), prod)
	// } else {
	// 	c.String(404, "Información del producto ¡No existe! \n")

	// }
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Sabrina!",
		})
	})

	//router.GET("/products", GetAll)
	gopher := router.Group("/products")
	{
		gopher.GET("", GetAll)
	}
	// router.GET("/products", func(ctx *gin.Context) {
	// 	read, err := os.ReadFile("./products.json")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	data := string(read)
	// 	ctx.JSON(200, gin.H{"info": data})
	// })

	router.Run()
}
