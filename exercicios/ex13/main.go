// Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.

package main

import (
	"fmt"
)

func main() {
	for x := 65; x <= 90; x++ {
		fmt.Println("----------")
		for i := 0; i < 3; i++ {
			fmt.Printf("%#U\n", x)
		}
	}
}
