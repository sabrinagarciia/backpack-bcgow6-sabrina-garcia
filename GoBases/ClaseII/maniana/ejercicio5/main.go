package main

import "fmt"

func foodDog(amount int) int {
	return amount * 10000
}

func foodCat(amount int) int {
	return amount * 5000
}

func foodHamster(amount int) int {
	return amount * 200
}

func foodTarantula(amount int) int {
	return amount * 150
}

func Animal(animal string) (func(int) int, string) {
	switch animal {
	case "dog":
		return foodDog, ""
	case "cat":
		return foodCat, ""
	case "hamster":
		return foodHamster, ""
	case "tarantula":
		return foodTarantula, ""
	default:
		return nil, "No, no, disculpá, ese animal no existe"
	}
}

func main() {
	animalDog, msg := Animal("dog")

	var amount float64
	amount += float64(animalDog(5))

	// imprimo msg porque sino golang me tira error diciendo que se declara pero no se usa
	fmt.Println(msg)

}
