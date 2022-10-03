package main

import (
	"fmt"
	"os"
)

type miError struct {
	status int
	msg string
}

func (e* miError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func miErrorTest (status int) (int, error) {
	if status >= 400 {
		return 500, &miError {
			status : status,
			msg: "algo salió mal",
		}
	}
	return 200, nil
}

func mainEjemplo() {
	status, err := miErrorTest( 300 )
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d, Funciona!", status)
}