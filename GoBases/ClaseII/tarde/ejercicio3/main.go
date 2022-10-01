package main

//Estructuras
type product struct {
	Type, Name	string
	Price		float64
}

type store struct {
	Products	[]product
}

//Metodos
func (p product) CalcularCosto() float64 {
	finalPrice := p.Price

	switch {
	case p.Type == "Pequeño":
		return finalPrice
	case p.Type == "Mediano":
		finalPrice = p.Price * 1.03
		return finalPrice
	case p.Type == "Grande":
		finalPrice = p.Price * 1.06
		return finalPrice + 2500.00
	default:
		return 0.0
	}
}

func (s store) total() float64 {
	totalPrice := 0.0
	for _, p := range s.Products {
		productPrice := p.CalcularCosto()
		totalPrice += productPrice
	}
	return totalPrice
}

func (s *store) agregar(product *product) {
	s.Products = append(s.Products, *product)
}

//Interfaces
type Product interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	total() float64
	agregar()
}

//Funciones
func nuevoProducto(productType, name *string, price *float64) *product {
	return &product{Type: *productType, Name: *name, Price: *price}
}

func nuevaTienda(products *[]product) *store {
	return &store{Products: *products}
}


func main() {

}