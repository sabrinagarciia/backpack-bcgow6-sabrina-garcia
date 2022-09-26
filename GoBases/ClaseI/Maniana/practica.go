package main

import (
	"fmt"
)

func main() {
	// ACTIVIDAD 1
	nombresito := "FRANCO PESENDA"
	direccion := "Altolaguirre"
	fmt.Println("Nombre: " + nombresito)
	fmt.Println("Direcciòn: " + direccion)
	//ACTIVIDAD 2
	temperatura := 25.3
	humedad := 17.5
	presion := 1025.5
	fmt.Println(
		"Temperatura: ", temperatura, "\n",
		"Humedad: ", humedad, "\n",
		"Presión: ", presion)
	//ACTIVIDAD 3
	var nombre string
	var apellido string
	var edad int
	apellidoDos := 6
	var licencoaDeComnducir = true
	var estaturaDeLaPersona int
	cantidadDeHijos := 2
	fmt.Println("Nombre: " + nombre)
	fmt.Println("Direcciòn: " + apellidoDos, cantidadDeHijos, licencoaDeComnducir)
	//ACTIVIDAD 4
	var apellidoA4 string = "Gomez"
	var edad1 int = 35
	boolean := "false"
	var sueldo float64 = 45857.90
	var nombre1 string = "Julian"
	fmt.Println("Nombre: " + nombre1 + apellidoA4)
	fmt.Println(apellidoDos, edad1, boolean, sueldo)
}