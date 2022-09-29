package main

import (
	//"fmt"
	"errors"
)

func averageStudentGrade(grades ...int) (int, error) {
	var result int
	var count int
	var sumGrades int
	for _, value := range grades {
		if value < 0 {
			err := errors.New("no pueden existir valores negativos")
			return 0, err
		} else {
			count += 1
			sumGrades += value
			result = sumGrades / count
		}
	}
	return result, nil
}

// func main() {
// 	fmt.Println(averageStudentGrade(3, -3, 3))
// }