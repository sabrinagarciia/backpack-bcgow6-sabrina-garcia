package main

import (
	//"errors"
	"fmt"
)

// type error interface {
// 	Errors()	string
// }

type MyCustomError struct {
	msg		string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf(e.msg)
}

func MyCustomErrorTest(salary int) (string, error) {
	if salary < 150000 {
		return "", &MyCustomError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 300

	fmt.Println(MyCustomErrorTest(salary))
}