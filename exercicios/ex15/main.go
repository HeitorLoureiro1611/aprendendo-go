// Crie um loop utilizando a sintaxe: for {}
// Utilize-o para demonstrar os anos desde que você nasceu.

package main

import (
	"fmt"
)

func main() {

	ano_nascimento := 2003
	ano_atual := 2026

	for {

		if ano_nascimento > ano_atual {
			break
		}

		fmt.Printf("%d\n", ano_nascimento)
		ano_nascimento++

	}

	fmt.Println("fim do loop")
}
