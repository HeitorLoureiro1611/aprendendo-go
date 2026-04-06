package main

import (
	"fmt"
)

// primeira declaração do tipo int
var x int = 10

// declaração do tipo string
var y string = "texto!"

// Inicialização do tipo float
var z float64

// z = 3.1415 -> isso seria um problema, pois está sendo atribuido fora de um escopo
// só se pode atribuir valores a variáveis quando estão em algum escopo
func main() {

	// modificação do tipo int, sem problemas
	x = 99

	// declaração do valor
	z = 3.1415

	// uso das variáveis
	fmt.Printf("%d -> %T\n", x, x)
	fmt.Printf("%s -> %T\n", y, y)
	fmt.Printf("%.4f -> %T\n", z, z)

}
