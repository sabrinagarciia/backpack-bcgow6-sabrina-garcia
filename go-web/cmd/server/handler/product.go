package handler

import(
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/svbri/backpack-bcgow6-sabrina-garcia/go-web/internal/products"
)
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