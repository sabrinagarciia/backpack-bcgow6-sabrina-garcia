package main

import (
	"fmt"
	"os"
)

func main() {
	salary := 150000
	err := errorSalaryTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) //salgo del programa con este método
	}
	fmt.Printf("Debe pagar impuesto")
}

type errorSalary struct {
	msg string
}

func (e* errorSalary) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func errorSalaryTest (salary int) (error) {
	if salary < 150000 {
		return &errorSalary {
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}
