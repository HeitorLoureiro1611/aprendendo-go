// Crie uma função que receba um parâmetro variádico do tipo int
// e retorne a soma de todos os ints recebidos;
// Passe um valor do tipo slice of int como argumento para a função;
// Crie outra função, esta deve receber um valor do tipo slice of int
// e retornar a soma de todos os elementos da slice;
// Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func somaVariatico(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}

	return total

}

func somaSlice(i []int) int {
	total := 0

	for _, v := range i {
		total += v
	}
	return total
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(somaVariatico(slice...))
	fmt.Println(somaSlice(slice))
}
