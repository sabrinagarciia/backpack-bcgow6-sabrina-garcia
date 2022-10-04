package main

import (
	"fmt"
	"net/http"
	//"os"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type Product struct {
	ID        int     `json:"id" binding:"required" validate:"required"`
	Name      string  `json:"name" binding:"required" validate:"required"`
	Color     string  `json:"color" binding:"required" validate:"required"`
	Price     float64 `json:"price" binding:"required" validate:"required"`
	Stock     int     `json:"stock" binding:"required" validate:"required"`
	Code      string  `json:"code" binding:"required" validate:"required"`
	Published bool    `json:"published" binding:"required" validate:"required"`
	Date      string  `json:"date" binding:"required" validate:"required"`
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
var validate *validator.Validate
func Save(c *gin.Context) {
	token := c.GetHeader("token")

	if token != "aaaaa" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
		return
	}

	var req Product

	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(404, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	validate = validator.New()
	err := validate.Struct(req)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
	//lastID++
	req.ID = len(products) + 1
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
