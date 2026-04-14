// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// Demonstre os valores do map utilizando range.
// Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

package main

import (
	"fmt"
)

type pessoa struct {
	nome       string
	sobrenome  string
	sorvetefav []string
}

func main() {
	// neste caso eu uso o make para criar o map usando um struct
	pessoas := make(map[string]pessoa)

	pessoas["Johnson"] = pessoa{
		nome:       "Jonatas",
		sobrenome:  "Johnson",
		sorvetefav: []string{"baunilha", "morango", "napolitano"},
	}

	pessoas["Januairo"] = pessoa{
		nome:       "Janaina",
		sobrenome:  "Januairo",
		sorvetefav: []string{"pistache", "ninho-trufado", "oreo"},
	}

	for _, v := range pessoas {
		fmt.Println("=============================================")
		fmt.Printf("nome: %s | sobrenome: %s", v.nome, v.sobrenome)
		fmt.Println("\nSabores favoritos de sorvete: ")
		for x := range v.sorvetefav {
			fmt.Println("-", v.sorvetefav[x])
		}
	}

}
