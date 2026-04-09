// Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
// Demonstre estes valores.

package main

import (
	"fmt"
)

func main() {
	a := 10 > 11
	b := 9 != 10
	c := 10 == 8
	d := 8 <= 9

	fmt.Printf("%v, %v, %v, %v\n", a, b, c, d)
}
