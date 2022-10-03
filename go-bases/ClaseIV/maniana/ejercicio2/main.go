package main

import (
	"errors"
	"fmt"
)

// type error interface {
// 	Errors()	string
// }

type MyCustomError struct {
	msg		error
}

// func (e *MyCustomError) Error() error {
// 	return 
// }

func MyCustomErrorTest(salary int) (string, error) {
	if salary < 1506000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 300

	fmt.Println(MyCustomErrorTest(salary))
}