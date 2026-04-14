// Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
// - Nome
// - Sobrenome
// - Sabores favoritos de sorvete
// Crie dois valores do tipo "pessoa" e demonstre estes valores,
// utilizando range na slice que contem os sabores de sorvete.

package main

import (
	"fmt"
)

type pessoa struct {
	nome       string
	sobrenome  string
	saboresfav []string
}

func main() {
	pessoa1 := pessoa{
		nome:       "Jonatas",
		sobrenome:  "Johnson",
		saboresfav: []string{"baunilha", "morango", "napolitano"},
	}
	pessoa2 := pessoa{
		nome:       "Janaina",
		sobrenome:  "Januairo",
		saboresfav: []string{"pistache", "ninho-trufado", "oreo"},
	}

	fmt.Printf("nome: %s | sobrenome: %s\nSabores favoritos de sorvete:\n", pessoa1.nome, pessoa1.sobrenome)
	for x := range pessoa1.saboresfav {
		fmt.Println(pessoa1.saboresfav[x])
	}
	fmt.Println("===================================")

	fmt.Printf("nome: %s | sobrenome: %s\nSabores favoritos de sorvete:\n", pessoa2.nome, pessoa2.sobrenome)
	for x := range pessoa2.saboresfav {
		fmt.Println(pessoa2.saboresfav[x])
	}
}
