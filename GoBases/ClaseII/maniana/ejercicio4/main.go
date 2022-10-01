package main

func minimum(grades ...int) int {
	minimum := 10
	for _, grade := range grades {
		if (grade < minimum) {
			minimum = grade
		}
	}
	return minimum
}

func maximum(grades ...int) int {
	maximum := 0
	for _, grade := range grades {
		if (grade > maximum) {
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

func operation(minmaxavg string) func(grades ...int) int {
	switch minmaxavg {
		case "minimum":
			return minimum
		case "maximum":
			return maximum
		case "average":
			return average
	}
	return nil
}

func main() {
	//minFunc, err := operation("minimum")

	 
}