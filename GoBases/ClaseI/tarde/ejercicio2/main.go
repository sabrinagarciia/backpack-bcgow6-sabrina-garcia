package main

import (
	"fmt"
)

func isValidForLoan(age int, yearsWorking int, salary int) {
	if (age >= 22 && yearsWorking > 1) {
		interest := "con intereses."
		if (salary > 100000) {
			interest = "y no tiene interés."
		}
		fmt.Println("Aplica para préstamo" + " " + interest)
	} else {
		fmt.Println("No puede acceder al préstamo.")
	}
}

func main() {
	isValidForLoan(23, 3, 120000)
}