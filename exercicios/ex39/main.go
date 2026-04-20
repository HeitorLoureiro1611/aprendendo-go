// Crie uma função que retorna uma função.
// Atribua a função retornada a uma variável.
// Chame a função retornada.

package main

import (
	"fmt"
)

func ondeestou() func() {
	return func() {
		fmt.Println("eu estou dentro de 2 funções!")
	}
}

func main() {

	variavel := ondeestou()

	variavel()

}
