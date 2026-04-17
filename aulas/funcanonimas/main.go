package main

import "fmt"

func main() {

	// funçoes lambda
	// funções que não vou precisar reutilizar senão nesse unico contexto, "descartável"
	// funções anonimas são "funções inline"
	// define ela sem dar um nome
	x10 := func(n int) int {
		return n * 10
	}
	// comum usar funções anônimas com defer para capturar
	// erros (recover) ou fechar recursos, pois elas permitem agrupar várias
	// ações de limpeza em um único bloco.
	fmt.Println(x10(1))
	fmt.Println(x10(x10(1)))

}
