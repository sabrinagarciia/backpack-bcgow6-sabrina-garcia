package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("Empezando...")
	read, err := os.ReadFile("customers.txt")
	if err != nil {
		panic(err)
	}
	data := string(read)
	fmt.Println(data)
}