// Demonstre o funcionamento de um closure.
// Ou seja: crie uma função que retorna outra função,
// onde esta outra função faz uso de uma variável alem de seu scope.

package main

import (
	"fmt"
)

// função que não recebe nada e retorna uma função do tipo int
func contador() func() int {
	x := 0
	// retornando uma função do tipo int
	return func() int {
		// funcionamento da função
		x++
		// retorno do tipo int
		return x
	}
}

func main() {

	a := contador()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println()

	b := contador()
	fmt.Println(b())

}
