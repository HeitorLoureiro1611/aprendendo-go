// Utilizando a solução do exercício anterior:

// 1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y".
// O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
// Na função main:
// 1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e,
// utilizando o operador "=", atribua o valor de "x" a "y"
// 2. Demonstre o valor de "y"
// 3. Demonstre o tipo de "y"

package main

import (
	"fmt"
)

type numero int

var x numero
var y int

func main() {

	fmt.Println("Valor de x: ", x)
	fmt.Println("Valor de y: ", y)

	fmt.Printf("tipo de x -> %T\n", x)
	fmt.Printf("tipo de y -> %T\n", y)

	x = 42
	y = int(x)

	fmt.Println("Valor de y: ", y)
	fmt.Println("Valor de x: ", x)
}
