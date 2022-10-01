package main

import (
	"errors"
	"fmt"
)

func minimum(grades ...int) int {
	minimum := 10
	for _, grade := range grades {
		if grade < minimum {
			minimum = grade
		}
	}
	return minimum
}

func maximum(grades ...int) int {
	maximum := 0
	for _, grade := range grades {
		if grade > maximum {
			maximum = grade
		}
	}
	return maximum
}

func average(grades ...int) int {
	count := 0
	sum := 0
	for _, grade := range grades {
		count += 1
		sum += grade
	}
	avg := sum / count
	return avg
}

func operation(minmaxavg string) (func(grades ...int) int, error) {
	switch minmaxavg {
	case "minimum":
		return minimum, nil
	case "maximum":
		return maximum, nil
	case "average":
		return average, nil
	}
	return nil, errors.New("algo salió mal che")
}

func main() {
	minFunc, err := operation("minimum")

	minValue := minFunc(5,6,8,4,2,9)

	fmt.Println(minValue)
	fmt.Println(err)
}
