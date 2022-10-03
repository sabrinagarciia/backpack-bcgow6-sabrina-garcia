package main

import (
	"fmt"
	//"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID			int
	Name		string
	Color		string
	Price		float64
	Stock		int
	Code		string
	Published	bool
	Date		string
}

/*func GetAll(/*w http.ResponseWriter, req *http.Request) {
	read, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err)
	}

	data := string(read)
	c.JSON(200, gin.H{"info": data})
}*/

func main() {
	// ------- Ejercicio 1 de la clase 1 -------
	router := gin.Default()

	router.GET("/hello", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "Hello Sabrina!",
		})
	})



	router.GET("/products", func(ctx *gin.Context) {
		read, err := os.ReadFile("./products.json")
		if err != nil {
			fmt.Println(err)
		}
	
		data := string(read)
		ctx.JSON(200, gin.H{"info": data})
	})
	router.Run()
}