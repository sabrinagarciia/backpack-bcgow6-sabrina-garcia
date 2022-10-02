package main

import (
	"fmt"
)

func MyCustomErrorTest(salary int) (string, error) {
	if salary < 1506000 {
		return "", fmt.Errorf("error: el mínimo imponible es de $150.000 y el salario ingresado es de: $%d", salary)
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 300

	fmt.Println(MyCustomErrorTest(salary))
}