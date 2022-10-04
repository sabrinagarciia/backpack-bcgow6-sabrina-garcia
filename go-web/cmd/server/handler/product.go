package handler

import(
	//"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sabrinagarciia/backpack-bcgow6-sabrina-garcia/go-web/internal/products"
)

type request struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	Date      string  `json:"date"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll(ctx *gin.Context) {
	// var filtered []Product
	// query := c.Query("name")
	// for _, p := range products {
	// 	if query == p.Name {
	// 		filtered = append(filtered, p)
	// 	}
	// }

	// c.JSON(http.StatusOK, filtered)

	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	p, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, p)
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

// func GetId(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.String(404, "Información del producto ¡No existe! \n")
// 	} else {
// 		c.String(200, "Información del producto %s, nombre: %s \n", c.Param("id"), id)
// 	}
// }

