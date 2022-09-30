package main

import (
	"fmt"
)

func whichMonthIsIt(monthNumber int) {
	var month string
	switch {
		case monthNumber == 1:
			month = "Enero"
		case monthNumber == 2:
			month = "Febrero"
		case monthNumber == 3:
			month = "Marzo"
		case monthNumber == 4:
			month = "Abril"
		case monthNumber == 5:
			month = "Mayo"
		case monthNumber == 6:
			month = "Junio"
		case monthNumber == 7:
			month = "Julio"
		case monthNumber == 8:
			month = "Agosto"
		case monthNumber == 9:
			month = "Septiembre"
		case monthNumber == 10:
			month = "Octubre"
		case monthNumber == 11:
			month = "Noviembre"
		case monthNumber == 12:
			month = "Diciembre"
		default:
			fmt.Println("No es un mes válido.")
	}

	fmt.Println(month)
}

func main() {
	whichMonthIsIt(9)
}