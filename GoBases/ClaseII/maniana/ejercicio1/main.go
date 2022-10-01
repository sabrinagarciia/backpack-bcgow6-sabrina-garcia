package main

import (
	"fmt"
)

func calculateTax(sueldo float32) float32 {
	var result float32
	if sueldo > 50000 {
		result = sueldo * 0.17
		fmt.Println("El impuesto de este sueldo es de: ", result)
	} else if sueldo > 150000 {
		result = sueldo * 0.27
		fmt.Println("El impuesto de este sueldo es de: ", result)
	}
	return result
}
