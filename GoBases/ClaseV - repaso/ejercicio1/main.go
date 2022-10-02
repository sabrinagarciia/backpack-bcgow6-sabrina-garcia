package main

import (
	"os"
	"fmt"
)

func main() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Ejecución finalizada")
		}
	}()

	fmt.Println("Empezando...")
	read, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("aaaa")
	}
	data := string(read)
	fmt.Println(data)
	
}