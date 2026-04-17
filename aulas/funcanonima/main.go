package main

import "fmt"

func main() {

	x := 94

	// função anonima
	// manipulação ráida de dados
	// será usada em goroutines
	func(x int) int {
		return x + 40
	}(x)
	// ela acaba aqui e não pode ser chamada novamente

	fmt.Println("O valor é: ", x)

}
