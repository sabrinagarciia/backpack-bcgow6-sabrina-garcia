package ejercicio1

import(
	"fmt"
)

func LetterQuantity(word string) {
	fmt.Printf("Número de letras: %d\n", len(word))
	for i := 0; i < len(word); i++ {
		//Esto imprime todas las letras en una misma línea
		//fmt.Printf("%s", string(word[i]))
		
		//Esto imprime todas las letras una por línea
		fmt.Println(string(word[i]))
	}
	fmt.Println()
}