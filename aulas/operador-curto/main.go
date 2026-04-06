package main

import (
	"fmt"
)

func main() {

	// Atribuição rápida e automatica com :=
	numero := 10 + 10
	texto := "Isso é uma string"
	boolean := false

	// fmt.println printa o que está entre parenteses e retorna 2 valores
	// quantidade de bytes e um valor de erro
	nbytes, erro := fmt.Println(numero, texto, boolean)

	// 28 bytes usados e nenhum erro encontrado <nil>
	fmt.Println(nbytes, erro)
}
