// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.

package main

import (
	"fmt"
)

const (
	anoatual = iota + 2026
	b
	c
	d
	e
)

func main() {

	fmt.Println("ano atual: ", anoatual)
	fmt.Println("Proximos 4 anos: ")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
