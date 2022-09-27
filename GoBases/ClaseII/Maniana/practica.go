package main

import (
	"fmt"
)


func devuelveImpuesto (sueldo float32) float32 {
	var resultado float32
	if sueldo > 50000 {
		resultado = sueldo * 0.17
		fmt.Println("El impuesto de este sueldo es de: ", resultado)
	} else if sueldo > 150000 {
		resultado = sueldo * 0.27
		fmt.Println("El impuesto de este sueldo es de: ", resultado)
	}
	return resultado
}
	
func main() {
	devuelveImpuesto(155000)
}
