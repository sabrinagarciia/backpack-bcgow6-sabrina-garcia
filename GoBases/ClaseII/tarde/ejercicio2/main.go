package main

import "fmt"

type Matrix struct {
	Elements  []float64
	Rows      int
	Columns   int
	Quadratic bool
	MaxValue  int
}

func (m *Matrix) Set(values ...float64) {
	m.Elements = values
}

func (m Matrix) Print() {
	for i := 0; i < m.Columns; i++ {
		for _, values := range m.Elements {
			fmt.Println(m.Elements)
		}
	}
}

func main() {
	m1 := Matrix{
		Rows:      8,
		Columns:   3,
		Quadratic: false,
		MaxValue:  3,
	}

	m1.Set(50000.00, 00.000010, 00.020, 10.00)

	m1.Print()
}
