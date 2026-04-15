package main

import (
	"fmt"
)

// toda função é *pass by value*!!!!
func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7}
	linha()

	// defer é uma declaração de ""deixe para ultima hora"" no momento da execução do programa
	// ou seja, será colocado como ultima coisa na minha função
	// Funcionamento de pilha quando é usado multiplos defer "primeiro a declarar, ultimo a sair"
	defer fmt.Println(intx2(9))

	// Pra conseguir usar o conteúdo de uma slice como argumento de uma função
	// é necessário utilizar o operador "..." indicando "use o conteúdo" | "Enumere"
	fmt.Println(soma(numeros...))

	// Por se tratar de uma função que recebe argumentos variáticos
	// eu posso receber nenhum argumento sem cair em um erro
	fmt.Println(soma())

}

// função básica
func linha() {
	fmt.Println("-------------------------------------------------------------------------------------------------")
}

// chamada | nome | (parametro tipo) | tipo de retorno
func intx2(x int) int {
	// retorno
	return x * 2
}

// Func com multiplos retornos
// argumento variático é que não tem um numero especifico de inputs daquele tipo
// deve ser sempre o ultimo a ser colocado nos parâmetros.
// Aqui x é considerado uma slice
func soma(x ...int) (int, int) {
	total := 0
	// desconsidere a chave, use apenas o valor dos numeros que estão em x
	for _, i := range x {
		total += i
	}
	return total, len(x)
}
