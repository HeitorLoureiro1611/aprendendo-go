package main

import (
	"fmt"
)

// declaração de array (numeros finitos de espaço)
var array = [5]int{0, 0, 15, 20, 30}

// declaração de slice (array com espaço dinâmico)
var slice = []int{1, 2, 3, 4, 5, 6, 7} //...

func main() {
	// declaração direta na posição do array
	array[0] = 1
	array[1] = 10

	// printando o array formatado usando o range
	for i := range array {
		fmt.Printf("Indice: %d -> valor: %d\n", i, array[i])
	}

	// append, iteração de um novo numero no slice
	slice = append(slice, 10)
	fmt.Println(slice)

	// função de retornar um int com o tamanho do array
	fmt.Printf("\n%T, %d", len(array), len(array))

}
