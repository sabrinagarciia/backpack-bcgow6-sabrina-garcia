package products

import (
	"fmt"
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
/*
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
}*/

var products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Save(id int, name, color string, price float64, stock int, code string, published bool, date string) (Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price float64, stock int, code string, published bool, date string) (Product, error)
}

type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Save(id int, name, color string, price float64, stock int, code string, published bool, date string) (Product, error) {
	p := Product{id, name, color, price, stock, code, published, date}
	products = append(products, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, name, color string, price float64, stock int, code string, published bool, date string) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	updated := false
	for i := range products {
		if products[i].ID == id {
			p.ID = id
			products[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
}