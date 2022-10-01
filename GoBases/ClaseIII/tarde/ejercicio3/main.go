package main


type Product struct {
	Name		string
	Price, Quantity	int
}

type Service struct {
	Name		string
	Price		float64
	Minutes		int
}

type Maintenance struct {
	Name		string
	Price		float64
}

func totalProducts(products []Product) int {
	sum := 0
	for _, p := range products {
		sum += p.Price * p.Quantity
	}
	return sum
}

func totalService(services []Service) {
	
}

func finalTotal() {

}

func totalMaintenance() {
	
}

func main() {

}