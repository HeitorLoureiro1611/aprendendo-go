package main

import (
	"fmt"
)

func fatorial(x int) int {
	// caso o numero for 0 pra n crashar
	if x == 0 {
		return 0
	}
	// se for 1 acaba
	if x == 1 {
		return 1
	}
	// recursividade é utilizar uma função no retorno dela mesma
	// chamar ela mesma quando ela terminar
	// ciclo
	return x * fatorial(x-1)
	// 5 * 4 * 3 * 2 * 1 = 120
}

func main() {
	numero := 5
	fmt.Println(fatorial(numero))
}
