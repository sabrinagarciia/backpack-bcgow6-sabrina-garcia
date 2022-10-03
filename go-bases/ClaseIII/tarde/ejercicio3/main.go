package main

import (
	"fmt"
)

type Product struct {
	Name		string
	Price		float64
	Quantity	int
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

func totalProducts(products *[]Product, c chan float64) {
	sum := 0.0
	for _, p := range *products {
		sum += p.Price * float64(p.Quantity)
	}
	fmt.Println("Total productos $", sum)
	c <- sum
	close(c)
}

func totalService(services *[]Service, c chan float64) {
	sum := 0.0
	for _, s := range *services {
		if s.Minutes < 30 {
			sum += s.Price * 30
		} else {
			sum += float64(s.Minutes) * s.Price
		}
	}
	fmt.Println("Total servicios $", sum)
	c <- sum
	close(c)
}

func totalMaintenance(maintenance *[]Maintenance, c chan float64) {
	sum := 0.0
	for _, m := range *maintenance {
		sum += m.Price
	}

	fmt.Println("Total mantenmiento $", sum)
	c <- sum
	close(c)
}

func main() {

}