package main

import "fmt"

func main() {
	// criando uma slice de strings
	slices := []string{"mamão", "papaya", "kiwi", "abacaxi", "limão"}

	fmt.Println("Primeira versão ->\t", slices)

	// retirando um item de uma slice:
	// append, primeiro arg: até onde ir, segundo arg: por onde começar
	// "use os itens que estão entre o inicio e o indice numero 1 (mamão)
	// e junte com os items que estão entre o indice numero 2 até o fim"
	// "..." -> unfurl
	// "..." no fim do append é pra indicar que eu quero pegar os valores dentro de slice[2:]
	slices = append(slices[:1], slices[2:]...)

	fmt.Println("Segunda versão ->\t", slices)

	// fatiando uma slice (pegando as informações seguindo os indices)
	fatia := slices[1:3]

	// printando a variável que guardou as informações fatiadas da slice
	fmt.Println("Versão fatiada ->\t", fatia)
}
