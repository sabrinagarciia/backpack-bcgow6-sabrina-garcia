package main

import (
	"fmt"
)

func salaryOperation(valuePerHour int, hours int, bonus int) int {
	result := valuePerHour * hours
	calculatedBonus := result * bonus / 100
	if bonus > 0 {
		return result + calculatedBonus
	}
	return result
}

func calculateSalary(hoursPerMonth int, category string) int {
	var salary int
	switch {
		case category == "c":
			//salary = 1000 * hoursPerMonth
			salary = salaryOperation(1000, hoursPerMonth, 0)
		case category == "b":
			//salary = 1500 * hoursPerMonth
			salary = salaryOperation(1500, hoursPerMonth, 20)
		case category == "a":
			salary = salaryOperation(3000, hoursPerMonth, 50)
		default:
			return 0
	}
	return salary
}

func main() {
	fmt.Println( calculateSalary(1000, "a"))
}