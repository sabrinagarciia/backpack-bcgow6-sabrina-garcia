package main

import (
	"log"
	"os"
)

func main() {
	data := []byte("1, $20, Alta\n2, $60, Baja\n")

	//err := os.WriteFile("./info.csv", []byte("Id, Precio, Calidad\n"), 0666)
	err := os.WriteFile("./info.csv", data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}