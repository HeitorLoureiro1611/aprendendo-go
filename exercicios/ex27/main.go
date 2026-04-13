// Crie um map com key tipo string e value tipo []string.
// Key deve conter nomes no formato sobrenome_nome
// Value deve conter os hobbies favoritos da pessoa
// Demonstre todos esses valores e seus índices.

package main

import (
	"fmt"
)

func main() {
	tabela := map[string]string{}

	tabela["Heitor_Loureiro"] = "Jogar"
	tabela["Thainá_Silva"] = "Ler"
	tabela["Gopher_Golang"] = "Programar"

	for k, v := range tabela {
		fmt.Println(k, "->", v)
	}

}
