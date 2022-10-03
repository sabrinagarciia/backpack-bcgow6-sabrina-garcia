package main

import "fmt"

type Students struct {
	Name		string
	Lastname	string
	DNI			string
	Date		string
}

func (s Students) detalle() {
	fmt.Println("Nombre: " + s.Name)
	fmt.Println("Apellido: " + s.Lastname)
	fmt.Println("DNI: " + s.DNI)
	fmt.Println("Fecha: " + s.Date)
}

func main() {
	s1 := Students{
		Name: "Rory",
		Lastname: "Gilmore",
		DNI: "00000000",
		Date: "30-09-2022",
	}

	s1.detalle()
}