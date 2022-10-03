package main

import(
	"strings"
	"fmt"
	"os"
)

func main() {
	read, err := os.ReadFile("../ejercicio1/info.csv")
	if err != nil {
		fmt.Println(err)
	}

	data := string(read)
	fmt.Println("ID\t\t\tPrecio\t\t\tCantidad")
	fmt.Println(strings.ReplaceAll(data, ",", "\t\t\t"))

}