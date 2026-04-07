package main

import (
	"fmt"
)

// inicialização do tipo booleano
var x bool

func main() {
	// valor de inicialização de um bool
	fmt.Println(x)

	x = true

	// valor do bool atribuido
	fmt.Println(x)
	// declarando um bool de uma expressão logica (5 é maior que 10? falso!) retorna false
	x = 5 > 10

	// false
	fmt.Println(x)

}
