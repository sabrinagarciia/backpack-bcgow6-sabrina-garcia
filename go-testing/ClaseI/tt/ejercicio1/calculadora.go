package testing

import "fmt"

func Restar(a, b int) (int, error) {
	if b > a {
		return 0, fmt.Errorf("b no puede ser mayor a a")
	}
	return a - b, nil
}