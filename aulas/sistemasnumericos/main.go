package main

import (
	"fmt"
)

// palavra reservada para contagem de numero inteiro não tipado em ordem (0,1,2,3...)
const x = iota + 1*100

func main() {

	// mostrando o valor em: decimal, binario e hexadecimal
	fmt.Printf("%d, %b, %#x \n", x, x, x)

}
