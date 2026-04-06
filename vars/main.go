package main

import (
	"fmt"
)

// var é utilizado para declarar uma variável em escopo global "package level scope"
// devo utilizar o operador =
var x = 10

func main() {

	y := 90
	// printo a variável global, mesmo que não tenha sido declarada nesse escopo
	fmt.Println(x)
	fmt.Println(y)

	// Aqui estou levando a variável global para outra função, que também pode ler ela
	dobro(x)
}

func dobro(num int) {
	fmt.Println(num * 2)
}
