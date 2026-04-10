// Crie um loop utilizando a sintaxe: for condition {}
// Utilize-o para demonstrar os anos desde que você nasceu.

package main

import (
	"fmt"
)

func main() {
	ano_nascimento := 2003
	ano_conta := 2026
	for ano_nascimento <= ano_conta {
		fmt.Println(ano_nascimento)
		ano_nascimento++
	}
}
