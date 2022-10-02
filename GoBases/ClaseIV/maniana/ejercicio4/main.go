package main

import (
	"errors"
	"fmt"
)

func monthlySalary(hours int, hourlyRate float64) (float64, error) {
	monthly := hourlyRate * float64(hours)
	if monthly >= 150000.00 {
		monthly = 150000.00 * 0.9

		return monthly, nil
	} else if monthly < 80000.00 || monthly < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	return monthly, nil
}

func bonus(salary float64, months int) (float64, error){
	//[mejor salario del semestre] / 12 * [meses trabajados en el semestre]
	if months < 0 {
		return 0, fmt.Errorf("El trabajador no tiene meses trabajados. Valor %d inválido", months)
	}

	bonus := salary / 12 * float64(months)
	return bonus, nil
}

func main() {
	
}