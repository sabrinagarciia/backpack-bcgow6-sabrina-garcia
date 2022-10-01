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
		return finalPrice + 2500
	default:
		return 0.0
	}
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

func nuevaTienda() /*Ecommerce*/ {
	
}


func main() {

}