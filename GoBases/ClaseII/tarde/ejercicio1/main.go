package main

import "fmt"

type Students struct {
	Name		string
	Lastname	string
	DNI			string
	Date		string
}

func detalle(student Students) {
	fmt.Println("Nombre: " + student.Name)
	fmt.Println("Apellido: " + student.Lastname)
	fmt.Println("DNI: " + student.DNI)
	fmt.Println("Fecha: " + student.Date)
}

func main() {
	s1 := Students{
		Name: "Rory",
		Lastname: "Gilmore",
		DNI: "00000000",
		Date: "30-09-2022",
	}

	detalle(s1)
}