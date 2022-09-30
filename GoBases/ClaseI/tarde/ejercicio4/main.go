package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{
		"Benjamin": 20, 
		"Nahuel": 26, 
		"Brenda": 19, 
		"Darío": 44, 
		"Pedro": 30,
	}

	// Cuento empleados mayores de 21
	var olderThan21 int

	for _, e := range employees {
		if e > 21 {
			olderThan21++
		}
	}

	fmt.Printf("%d empleados son mayores de 21 años\n", olderThan21)

	// Agrego un empleado
	employees["Federico"] = 25
	fmt.Println(employees)

	// Elimino un empleado
	delete(employees, "Pedro")
	fmt.Println(employees)
}