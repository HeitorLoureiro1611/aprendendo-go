// Crie uma função que retorne um int
// Crie outra função que retorne um int e uma string
// Chame as duas funções
// Demonstre seus resultados

package main

import "fmt"

func retint(i int) int {
	return i * 10
}

func retStr(s string) string {
	return s + "eba!"
}

func main() {
	fmt.Println(retStr("ola!"))
	fmt.Println(retint(5))
}
