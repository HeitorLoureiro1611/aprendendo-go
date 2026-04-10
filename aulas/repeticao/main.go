package main

import (
	"fmt"
)

// não existe while...

func main() {
	// estrutura de for tradicional
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d -> %#U\n", i, i)
	}

	// for aninhado
	for row := 0; row < 3; row++ {
		fmt.Println()
		for col := 1; col <= 3; col++ {
			fmt.Printf("%d ", col)
		}
	}

	// for com condição de parada (while)
	x := 0
	for x < 10 {
		fmt.Println("Oi")
		x++
	}

	y := 0
	// while true
	for {
		if y == 5 {
			break
		}
		fmt.Println("isso não acaba")
		y++
	}

}
