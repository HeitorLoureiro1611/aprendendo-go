package main

import (
	"fmt"
)

func main() {
	x := 10

	// condicional
	if !(x > 100) {
		fmt.Println("Oi")
	} else if x < 9 {
		fmt.Println("Tudo bem?")
	} else {
		fmt.Println("Tchau!")
	}

	//switch
	// if com varios casos diferentes
	y := 4
	switch {
	case y < 5:
		fmt.Println("eita, menor que 5")
	case y > 5:
		fmt.Println("MAIOR que 5")
	case y == 5:
		fmt.Println("igual a 5")
	}

	for i := 6; i >= 1; i-- {
		switch {
		// , "ou"
		case i == 5, i == 6:
			fmt.Println("menor que 5 ou menor que 6")
		case i == 4:
			fmt.Println("menor que 4")
		case i == 3:
			fmt.Println("menor que 3")
		case i == 2:
			fmt.Println("menor que 2")
			// cai pro proximo case depois da execução do case atual
			fallthrough
		case i == 1:
			fmt.Println("menor que 1")
		// caso não entre em nenhum parametro dos cases
		default:
			fmt.Println("O que é isso?")
		}

	}

}
