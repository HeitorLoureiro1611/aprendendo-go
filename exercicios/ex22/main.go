// Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
// Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
// Do quinto ao último item do slice (incluindo o último item!)
// Do segundo ao sétimo item do slice (incluindo o sétimo item!)
// Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
// Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range numeros {
		fmt.Println(i, "-", v)
	}
	fmt.Printf("%T\n", numeros)

	fmt.Println("=======================")

	// respostas:
	fmt.Println(numeros[:3])
	fmt.Println(numeros[4:])
	fmt.Println(numeros[1:7])
	fmt.Println(numeros[2:9])
	fmt.Println(numeros[2 : len(numeros)-1])
}
